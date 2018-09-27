package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"eglass.com/utils"
	"github.com/go-redis/redis"
)

// Staff 唯一标识
type Staff uint32

// User 公众号用户标识
type User string

// StaffGroup {staff_id ...} group of staff
type StaffGroup map[Staff]bool

// StaffClients {staff_id -> client ...}
type StaffClients map[Staff]*Client

// Message from ws
type Message struct {
	client  *Client
	message []byte
}

// ChatMeta 由一个公众号用户和 多个staff 组成, 可持久化和加载
type ChatMeta struct {
	ID     uint32
	Name   string
	user   User
	staffs StaffGroup
	sTime  time.Time
}

// ChatContext 当前的一个公众号用户和 多个staff 组成
type ChatContext struct {
	meta        *ChatMeta
	user        User
	staffs      StaffClients
	bindStaffID float64
	status      int8
}

// StaffSendError send error
type StaffSendError struct {
	errType string
	err     error
}

// UserChats key公众号用户, value 为对接的staffs client
type UserChats map[User]*ChatContext

// StaffChats staff to {ChatContext ...}
type StaffChats map[Staff]map[*ChatContext]bool

// Exchanger 消息交换中心
type Exchanger struct {
	rwlock sync.RWMutex
	// 所有的会话meta信息
	chatMetas map[uint32]*ChatMeta
	// 当前所有会话
	chats map[*ChatContext]bool
	// user -> chats
	userChats UserChats
	// staff -> chats
	staffChats StaffChats
	// Registered clients.
	clients map[*Client]bool

	// staff id -> client
	staffs StaffClients

	// appid 下所有staff clients
	appidStaffs map[string]StaffClients

	// Inbound messages from the clients.
	messages chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
	// 记录已发送模板消息的staff
	staffSentTemplates map[Staff]bool

	// build map for staff -> client
	bindStaff    chan *Client
	userMessages chan []byte
	userBind     chan []byte
}

const defaultUser = User("default")

var (
	errWeixin   = errors.New("weixin result error")
	req         = utils.InitRequest()
	isProd      = os.Getenv("prod") == "true"
	db          = utils.GetDB(isProd)
	redisClient = utils.NewRedisClient(isProd)
	// TwoDays  缓存时间
	TwoDays = 2 * time.Hour * 24
)

func newExchanger() *Exchanger {
	ex := &Exchanger{
		messages:           make(chan Message),
		register:           make(chan *Client, 4),
		unregister:         make(chan *Client, 2),
		bindStaff:          make(chan *Client),
		clients:            make(map[*Client]bool),
		staffs:             make(StaffClients),
		chatMetas:          make(map[uint32]*ChatMeta),
		chats:              make(map[*ChatContext]bool),
		userChats:          make(UserChats),
		staffChats:         make(StaffChats),
		appidStaffs:        make(map[string]StaffClients),
		userMessages:       make(chan []byte, 10),
		userBind:           make(chan []byte),
		staffSentTemplates: make(map[Staff]bool),
	}
	ex.initDefaultChat()
	return ex
}

func (ex *Exchanger) initDefaultChat() {
	// create meta
	defaultUser := User("default")
	defaultChatMeta := ChatMeta{
		ID:     0,
		Name:   "use for internal test",
		user:   defaultUser,
		staffs: make(StaffGroup),
	}
	ex.chatMetas[defaultChatMeta.ID] = &defaultChatMeta
	// create context
	cc := ChatContext{
		meta:   &defaultChatMeta,
		user:   defaultUser,
		staffs: make(StaffClients),
	}
	ex.chats[&cc] = true
	ex.userChats[defaultUser] = &cc

}
func (ex *Exchanger) run() {
	for {
		select {
		case client := <-ex.register:
			ex.clients[client] = true
		case client := <-ex.unregister:
			if _, ok := ex.clients[client]; ok {
				staff := Staff(client.ID)
				delete(ex.clients, client)
				delete(ex.staffs, staff)
				appidStaffs, exist := ex.appidStaffs[client.Appid]
				if exist {
					delete(appidStaffs, staff)
				}
				close(client.send)
			}
		// 员工建立连接
		case client := <-ex.bindStaff:
			// staff -> client
			staff := Staff(client.ID)
			ex.staffs[staff] = client
			// 加入到 appid 营业员在线列表
			appidStaffs, exist := ex.appidStaffs[client.Appid]
			if !exist {
				// appid 不存在
				ex.appidStaffs[client.Appid] = StaffClients{staff: client}
			} else {
				appidStaffs[staff] = client
			}
			log.Printf("all staff : %v", ex.appidStaffs)
		case um := <-ex.userMessages:
			// 用户消息过来，分配在线的营业员 建立会话
			var userMsg map[string]interface{}
			error := json.Unmarshal(um, &userMsg)
			if error != nil {
				log.Printf("user message json parse error: %v \n", error)
			} else {
				ex.handleUserMessage(um, userMsg)
			}
		case message := <-ex.userBind:
			ex.handleBind(message)
		}
	}
}
func (ex *Exchanger) handleBind(message []byte) {
	log.Printf("receive bind msg: %v \n", string(message))
	var bindMsg map[string]interface{}
	error := json.Unmarshal(message, &bindMsg)
	if error != nil {
		log.Printf("user message json parse error: %v \n", error)
		return
	}
	// 推送绑定前的离线消息
	openid := bindMsg["openid"].(string)
	msgKey := fmt.Sprintf("chat-user-offine-%s", openid)
	offlineExist, _ := redisClient.Exists(msgKey).Result()
	if offlineExist == 0 {
		return
	}
	appid := bindMsg["appid"].(string)
	bindStaffID := bindMsg["bindStaffId"].(float64)
	// 更新缓存
	staffUsers := fmt.Sprintf("chat-user-peer-staffs-%d", Staff(bindStaffID))
	redisClient.SAdd(staffUsers, openid).Result()
	// user hisotry msgs -> bind staff
	msgs, _ := redisClient.LRange(msgKey, 0, -1).Result()
	redisClient.Expire(msgKey, 0).Result()
	User := User(openid)
	cc, exist := ex.userChats[User]
	if !exist {
		cc = &ChatContext{
			user:        User,
			staffs:      ex.appidStaffs[appid],
			bindStaffID: bindStaffID,
		}
		ex.chats[cc] = true
		ex.userChats[User] = cc
	}
	_, staffExist := cc.staffs[Staff(bindStaffID)]
	// 离线
	if !staffExist {
		length := len(msgs)
		if length == 0 {
			return
		}
		StoreStaffOfflineMessage(int64(bindStaffID), openid, msgs[length-1], int64(length))
	}

}

// 用户发过来消息
func (ex *Exchanger) handleUserMessage(raw []byte, userMsg map[string]interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Print(r)
		}
	}()
	u := userMsg["openid"].(string)
	User := User(u)
	Appid := userMsg["appid"].(string)
	// store权限的员工
	staffIds := userMsg["staffIds"].([]interface{})
	thirdId := int64(userMsg["third_id"].(float64))
	name, _ := userMsg["name"].(string)
	// 会话是否存在
	bindStaffID := userMsg["bindStaffId"].(float64)
	cc, exist := ex.userChats[User]
	if !exist {
		cc = &ChatContext{
			user:        User,
			staffs:      ex.appidStaffs[Appid],
			bindStaffID: bindStaffID,
		}
		ex.chats[cc] = true
		ex.userChats[User] = cc
	}
	cc.staffs = ex.appidStaffs[Appid]
	cc.bindStaffID = bindStaffID
	var chatStr string
	if bindStaffID > 0 {
		chatStr = "%3F" + fmt.Sprintf("id=%d&username=%s&chat_id=%d&openid=%s", int64(userMsg["id"].(float64)),
			url.QueryEscape(userMsg["name"].(string)), int64(bindStaffID), u)
	}
	// 所有在线的store权限的营业员
	go func() {
		for _, clientID := range staffIds {
			var id int64
			switch v := clientID.(type) {
			default:
				id = int64(v.(float64))
			case string:
				parsed, _ := strconv.ParseInt(clientID.(string), 10, 64)
				id = parsed
			}
			client, exist := cc.staffs[Staff(id)]
			// 发送给所有 staffIds
			if exist {
				// 在线直接发送
				client.send <- raw
			} else {
				ex.messageToStaff(id, thirdId, int64(bindStaffID), u, name, chatStr, raw)
			}
		}
	}()

}

// 发送消息到staff websocket
func (ex *Exchanger) messageToStaff(id, thirdID, bindStaffID int64, openid, name, chatStr string, raw []byte) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("send template fail %d:  %v", id, r)
		}
	}()
	// 不再保存 StoreStaffOfflineMessage(id, openid, string(raw), 1)
	// 5mins发送一次
	msgType := "customermsg"
	if bindStaffID > 0 {
		msgType = "msgnotify"
	}
	if ex.checkStaffIsSent(id) {
		return
	}
	ex.rwlock.Lock()
	_, exist := ex.staffSentTemplates[Staff(id)]
	if exist {
		ex.rwlock.Unlock()
		return
	}
	ex.staffSentTemplates[Staff(id)] = true
	ex.rwlock.Unlock()
	error := PushTemplateMsgToStaff(id, thirdID, name, chatStr, msgType)
	if error != nil {
		log.Printf("send to staff fail id: %d, reason: %s", id, error.Error())
	}
	time.AfterFunc(5*time.Minute, func() {
		ex.rwlock.Lock()
		delete(ex.staffSentTemplates, Staff(id))
		ex.rwlock.Unlock()
	})
}
func (ex *Exchanger) checkStaffIsSent(id int64) bool {
	ex.rwlock.RLock()
	_, exist := ex.staffSentTemplates[Staff(id)]
	ex.rwlock.RUnlock()
	return exist
}

func getContent(payload map[string]interface{}) string {
	content, exist := payload["content"]
	if exist {
		return content.(string)
	}
	mediaID, exist := payload["url"]
	if exist {
		return mediaID.(string)
	}
	mID, exist := payload["material_id"]
	if exist {
		return mID.(string)
	}
	card, exist := payload["name"]
	if exist {
		return card.(string)
	}
	return ""
}

// SendKfMessage 给公众号用户回复消息
func SendKfMessage(appid, openid, msgType string, thirdID, storeID, userID, staffID int32, payload map[string]interface{}) (int64, StaffSendError) {
	tx, error := db.DB.Begin()
	if error != nil {
		log.Fatalf("get transaction fail %v", error)
		tx.Rollback()
		return 0, StaffSendError{errType: "sql", err: error}
	}
	var sessionID int32
	tx.QueryRow("select id from e_chat_session where appid =? and openid = ? order by id desc limit 1 ", appid, openid).Scan(&sessionID)
	content := getContent(payload)
	duration := 0
	d, exist := payload["duration"]
	if exist {
		duration = int(d.(float64))
	}
	res, error := tx.Exec(`insert into e_chat_message(
	appid, openid, msg_type, third_id, store_id, user_id, from_staff, staff_id, session_id, send_time, content, duration)values
	(?,?,?,?,?,?,?,?,?, UNIX_TIMESTAMP(), ?, ?)`,
		appid, openid, msgType, thirdID, storeID, userID, 1, staffID, sessionID, content, duration)
	if error != nil {
		log.Fatalf("insert error %v", error)
		tx.Rollback()
		return 0, StaffSendError{errType: "sql", err: errors.New("message insert error")}
	}
	id, error := res.LastInsertId()
	if error != nil {
		log.Fatalf("insert id get error %v", error)
		tx.Rollback()
		return 0, StaffSendError{errType: "sql", err: errors.New("get message id error")}
	}
	payload["id"] = id
	result, error := req.KfSendMessage(appid, string(openid), msgType, payload)
	if error != nil {
		log.Printf("send kf message fail. message: %v, error: %v \n", payload, result)
		tx.Rollback()
		return 0, StaffSendError{errType: "network", err: error}
	}
	errCode := result["errcode"].(float64)
	if errCode > 0 {
		log.Printf("send kf message fail. message: %v, error: %v \n", payload, result)
		tx.Rollback()
		m := map[float64]string{
			40001: "公众号授权错误",
			40007: "媒体文件错误",
			45002: "文本消息过长(不超过682中文字符)",
			45015: "该顾客已取消关注，无法发送信息",
			45047: "发送条数超过限制(连续发送不能超过20条)",
		}
		return 0, StaffSendError{errType: "weixin", err: errors.New(m[errCode])}
	}
	tx.Commit()
	return id, StaffSendError{}
}

func (ex *Exchanger) handleStaffMessage(message []byte) (int64, StaffSendError) {
	log.Printf("receive msg: %v \n ", string(message))
	// 获取绑定 默认user contex
	var msgTo map[string]interface{}
	error := json.Unmarshal(message, &msgTo)
	if error != nil {
		log.Printf("message json parse error : %v ", error)
		return 0, StaffSendError{errType: "json", err: errors.New("message json parse error")}
	}
	openid := msgTo["openid"].(string)
	toUser := User(openid)
	bindStaffID := msgTo["staff_id"].(float64)
	appid := msgTo["appid"].(string)
	cc, exist := ex.userChats[toUser]
	if !exist {
		cc = &ChatContext{
			user:        toUser,
			staffs:      ex.appidStaffs[appid],
			bindStaffID: bindStaffID,
		}
		ex.chats[cc] = true
		ex.userChats[toUser] = cc
		log.Println(cc)
	}
	// send back to user
	msgType := msgTo["msg_type"].(string)
	userID := msgTo["user_id"].(float64)
	thirdID := msgTo["third_id"].(float64)
	storeID := msgTo["store_id"].(float64)
	payload := msgTo[msgType].(map[string]interface{})
	log.Printf("msg type: %s, %v\n", msgType, msgTo[msgType])
	// send to user
	id, staffSendError := SendKfMessage(appid, string(toUser), msgType, int32(thirdID), int32(storeID), int32(userID), int32(bindStaffID), payload)
	if staffSendError.errType != "" {
		log.Printf("kf send message fail \n: %v", staffSendError.err)
		return 0, staffSendError
	}
	go func() {
		db.DB.Exec(`update e_chat_session set last_active = ?, last_msg_id=? where openid = ? `, time.Now().Unix(), id, string(toUser))
	}()
	return id, StaffSendError{}
}

// StoreStaffOfflineMessage 离线存到redis
func StoreStaffOfflineMessage(staffID int64, openid string, lastMsg string, length int64) {
	// 离线存到redis
	staffUserStats := fmt.Sprintf("chat-staff-offine-%d-%s", staffID, openid)
	su, err := redisClient.Exists(staffUserStats).Result()
	if su == 0 || err == redis.Nil {
		redisClient.HMSet(staffUserStats, map[string]interface{}{"count": length, "lastMsg": lastMsg}).Err()
		redisClient.Expire(staffUserStats, TwoDays).Result()
	} else if su > 0 {
		redisClient.HIncrBy(staffUserStats, "count", length).Result()
		redisClient.HSet(staffUserStats, "lastMsg", lastMsg).Result()
		redisClient.Expire(staffUserStats, TwoDays).Result()
	}
	// staff -> ...openid
	staffUsers := fmt.Sprintf("chat-staff-offine-stats-%d", staffID)
	redisClient.SAdd(staffUsers, openid).Result()
	redisClient.Expire(staffUsers, TwoDays).Result()
}

// StoreUserOfflineMessage 离线存到redis
func StoreUserOfflineMessage(openid string, raw []byte) {
	// 离线存到redis
	msgKey := fmt.Sprintf("chat-user-offine-%s", openid)
	_, err := redisClient.Exists(msgKey).Result()
	error := redisClient.RPush(msgKey, string(raw)).Err()
	if err == redis.Nil {
		log.Printf("redis chat-staff not exist ")
	}
	if error != nil {
		log.Println("redis rpush fail", error)
	}
	redisClient.Expire(msgKey, time.Hour*24*2).Result()
}

func handleBindStaffMessage(bindStaffID float64, openid string, raw []byte, cc *ChatContext) {
	// 接入的staff
	// 绑定了staffid
	if bindStaffID == 0 {
		// 未绑定条件下, 只暂存用户消息
		StoreUserOfflineMessage(openid, raw)
		return
	}
	client, exist := cc.staffs[Staff(bindStaffID)]
	if !exist {
		// 不在线，用户消息放到此staff邮箱
		StoreStaffOfflineMessage(int64(bindStaffID), openid, string(raw), 1)
		return
	}
	// 在线直接发送
	client.send <- []byte(raw)
}

// PushTemplateMsgToStaff 推送离线模板消息  type "msgnotify" 未接入  "customermsg" 已接入
func PushTemplateMsgToStaff(id, thirdID int64, name, chatStr string, templateType string) (err error) {
	we, appid := 5, "wx7cc34ee2a43fe772"
	if isProd {
		we = 201
		appid = "wx5db40e19a0bd5244"
	}
	openid, staffName := "", ""
	err = db.DB.QueryRow(`select u.wx_id, s.name from e_staff s 
		join e_optometry_user u 
		on s.unionid = u.unionid
		where s.id = ? and s.unionid is not null and s.chat_notify_enabled = 1
		and s.third_id = ? and u.third_id = ? and u.wx_id is not null
		`, id, thirdID, we).Scan(&openid, &staffName)
	if err != nil {
		return nil
	}
	if openid == "" {
		log.Printf("请确认关注【营销互动专家】并打开消息接收权限: %d", id)
		return errors.New("no openid")
	}
	content, templateID, miniProgram := "", sql.NullString{}, sql.NullString{}
	err = db.DB.QueryRow("select content_list, msg_dev_id, miniprogram from e_notify_template where type = ? and third_id = ? ",
		templateType, we).Scan(&content, &templateID, &miniProgram)
	if err != nil {
		return err
	}
	if templateID.String == "" {
		return errors.New("no template id set ")
	}
	if miniProgram.String == "" {
		return errors.New("no miniapp set")

	}
	// 加载模板内容
	t := map[string]interface{}{}
	err = json.Unmarshal([]byte(content), &t)
	if err != nil {
		return err
	}

	for k := range t {
		temp := t[k].(map[string]interface{})
		temp["value"] = temp["content"]
		if k == "first" {
			temp["value"] = strings.Replace(temp["value"].(string), "name", staffName, -1)
		}
		if k == "keyword2" {
			temp["value"] = utils.ToFullTimeString(nil)
		}
		if k == "keyword1" {
			temp["value"] = name
		}
		t[k] = temp
	}
	// 加载小程序设置
	mini := map[string]interface{}{}
	err = json.Unmarshal([]byte(miniProgram.String), &mini)
	if err != nil {
		return err
	}
	mini["pagepath"] = mini["pagepath"].(string) + chatStr
	log.Printf("minipath: %s", mini["pagepath"])

	// send body
	Body := utils.Body(t)
	_, err = req.SendTemplateMsg(appid, openid, templateID.String, Body, "", mini)
	return err
}
