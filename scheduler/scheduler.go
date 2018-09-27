package main

import (
	"encoding/json"
	"log"
	"os"
	"sync/atomic"
	"time"

	"eglass.com/utils"
	"github.com/go-redis/redis"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Task 任务需要实现的接口
type Task interface {
	DueTimestamp() int64
	IsCycle() bool
	getCycle() int64
	Run(mysql sqlbuilder.Database) error
	GetID() int
}

type Scheduler interface {
	BulkImportTasks()
	AddTask(t Task)
	CancelTask(int)
	Schedule()
}

// TimerScheduler 基于timer实现的调度群
type TimerScheduler struct {
	tasks      map[int]*time.Timer
	totalTasks int32
	Mysql      sqlbuilder.Database
	redis      *redis.Client
}

// BulkImportTasks 批量导入任务
func (s *TimerScheduler) BulkImportTasks() {
	var tasks []ERemindRecord
	s.Mysql.Select("id", "remind_time").From("e_remind_record").Where("status = 'pending' and CURRENT_TIMESTAMP() < remind_time").All(&tasks)
	for _, task := range tasks {
		s.AddTask(task, 0)
	}
}

// AddTask 增加一个任务, 相同id会被覆盖
func (s *TimerScheduler) AddTask(t Task, delay int64) {
	afterInSeconds := t.DueTimestamp() + delay - time.Now().Unix()
	if afterInSeconds < 0 {
		log.Println("can not add task before current time")
		return
	}
	s.CancelTask(t.GetID())
	tm := time.AfterFunc(time.Second*time.Duration(afterInSeconds), func() {
		log.Printf("time arrive, run %d, %s", t.GetID(), utils.ToFullTimeString(nil))
		t.Run(s.Mysql)
		atomic.AddInt32(&s.totalTasks, -1)
		if t.IsCycle() {
			s.AddTask(t, t.getCycle())
		}
	})
	s.tasks[t.GetID()] = tm
	atomic.AddInt32(&s.totalTasks, 1)
	log.Printf("add task %d ", t.GetID())
}

// CancelTask 取消任务
func (s *TimerScheduler) CancelTask(id int) {
	timer, exist := s.tasks[id]
	if exist {
		timer.Stop()
		delete(s.tasks, id)
		log.Printf("cancel task: %d", id)
	}
}

// Schedule 开始调度
type RemindMessage struct {
	ID int `json:"id"`
}

func (s *TimerScheduler) HandleTask(ID int) {
	var m ERemindRecord
	s.Mysql.SelectFrom("e_remind_record").Where("id=?", ID).One(&m)
	s.AddTask(m, 0)
}
func (s *TimerScheduler) Schedule() {
	s.BulkImportTasks()
	remindChan := s.redis.Subscribe("remind").Channel()
	for {
		select {
		case r := <-remindChan:
			data := r.Payload
			var m RemindMessage
			error := json.Unmarshal([]byte(data), &m)
			if error != nil {
				log.Printf("invalid json format: %s", data)
				return
			}
			log.Printf("receive task to schedule: %d \n", m.ID)
			s.HandleTask(m.ID)
		}
	}
}

var isProd = false
var mysqlLogging = true

func main() {
	if os.Getenv("logging") == "false" {
		mysqlLogging = false
	}
	if env := os.Getenv("prod"); env != "" {
		log.Printf("prod: %s ", env)
		isProd = true
	}
	sess, error := utils.NewMysql(isProd, mysqlLogging)
	if error != nil {
		panic(error)
	}
	defer sess.Close()
	scheduler := &TimerScheduler{
		Mysql: sess,
		tasks: make(map[int]*time.Timer),
		redis: utils.NewRedisClient(false),
	}
	scheduler.Schedule()
}
