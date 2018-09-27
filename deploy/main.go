package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"

	"eglass.com/sta"
)

// WebApp static
var WebApp = map[string]bool{
	"eglass-erp":        true,
	"eglass":            true,
	"eglass-erp-mobile": true,
	"process":           true,
}

// AppDelayRunner all apps
type AppDelayRunner struct {
	apps       map[string]struct{}
	startTime  string
	hasPending bool
	// wg         sync.WaitGroup
}

// New get an AppDelayRunner
func New() *AppDelayRunner {
	return &AppDelayRunner{
		apps:       make(map[string]struct{}),
		startTime:  "12",
		hasPending: false,
	}
}
func (a *AppDelayRunner) runCmd(cmd string, args ...string) {
	// defer a.wg.Done()
	cmdExec := exec.Command(cmd, args...)
	stdOutError, _ := cmdExec.CombinedOutput()
	log.Printf("%s\n", stdOutError)
}

// 部署web目录
func (a *AppDelayRunner) deployWebFromTemp(app, tmp string) error {
	cmd := fmt.Sprintf("cd /epj/www && mv %s %s ", app, fmt.Sprintf("%s-old", app))
	if error := exec.Command("sh", "-c", cmd).Run(); error != nil {
		log.Printf("move current app failed: %v ", error)
		return error
	}
	mvCmd := fmt.Sprintf("cd /epj/www && mv %s/%s %s ", tmp, app, app)
	error := exec.Command("sh", "-c", mvCmd).Run()
	return error
}

func (a *AppDelayRunner) deployWebApp(app string) error {
	tmpName := fmt.Sprintf("temp-%s-%d", app, rand.Int31n(100000))
	clear := func() {
		exec.Command("sh", "-c", fmt.Sprintf("rm -rf %s", tmpName))
	}
	cmd := fmt.Sprintf("cd /epj/www && mkdir -p %s && scp -r root@docker.epeijing.cn:/epj/www/%s %s/", tmpName, app, tmpName)
	error := exec.Command("sh", "-c", cmd).Run()
	if error != nil {
		log.Printf("tmp dir create fail: %s, %v", cmd, error)
		clear()
		return error
	}
	resp, error := http.Get("http://main:8081/" + app)
	if error != nil || resp.StatusCode != 200 {
		log.Printf("request to main fail ")
		clear()
		return error
	}
	defer resp.Body.Close()
	var filesStats sta.FilesStats

	if error := json.NewDecoder(resp.Body).Decode(&filesStats); error != nil {
		log.Printf("invalid resp: %v", resp.Body)
		clear()
		return error
	}
	filesStatsLocal, error := sta.Scan(fmt.Sprintf("/%s/%s", tmpName, app))
	if error != nil {
		log.Printf("local dir read errror: %v ", error)
		clear()
		return error
	}
	if !reflect.DeepEqual(filesStats, filesStatsLocal) {
		return errors.New(fmt.Sprintf("dismatch: remote: %v, local: %v", filesStats, filesStatsLocal))
	}
	return a.deployWebFromTemp(app, tmpName)
}

// run all @
func (a *AppDelayRunner) runTask() {
	log.Println("run ssh task start")
	log.Println(a.apps)
	for k := range a.apps {
		_, exist := WebApp[k]
		if exist {
			go func(app string) {
				error := a.deployWebApp(app)
				if error != nil {
					log.Printf("deploy %s failed: %v ", app, error)
				} else {
					log.Printf("deploy web app %s ok ", app)
				}
			}(k)
		}
	}
	a.runCmd("sh", "-c", "cd /root/nginx-and-services && docker-compose pull && docker-compose up -d")
}

func (a *AppDelayRunner) pullRunApp(app string) {
	a.runCmd("sh", "-c", fmt.Sprintf("cd /root/nginx-and-services && docker pull docker.epeijing.cn:5000/%s && docker-compose up -d ", app))
}

func (a *AppDelayRunner) addPending(app string) {
	a.apps[app] = struct{}{}
	if a.hasPending {
		log.Println("has pending tasks")
		return
	}
	a.hasPending = true
	pendingTime := time.Hour * time.Duration(24-time.Now().Hour())
	go func(d time.Duration) {
		select {
		case <-time.After(d):
			log.Println("time arrived")
			a.runTask()
			a.hasPending = false
			a.apps = make(map[string]struct{})
		}
	}(pendingTime)
}
func (a *AppDelayRunner) runRealTime(app string) {
	_, exist := WebApp[app]
	if exist {
		go func() {
			error := a.deployWebApp(app)
			if error != nil {
				log.Printf("deploy %s failed: %v ", app, error)
			} else {
				log.Printf("deploy web app %s ok ", app)
			}
		}()
		return
	}
	go func() {
		time.Sleep(time.Minute * 2)
		a.pullRunApp(app)
	}()
}
func (a *AppDelayRunner) startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app := r.URL.Path[1:]
		if app == "" {
			fmt.Fprintf(w, "ok")
			return
		}
		if strings.Contains(app, "cancel") {
			app = strings.Split(app, "/")[1]
			delete(a.apps, app)
			log.Printf("%s removed", app)
			log.Println(a.apps)
			fmt.Fprintf(w, "%s removed", app)
			return
		}
		if strings.Contains(app, "realtime") {
			app = strings.Split(app, "/")[1]
			delete(a.apps, app)
			log.Printf("run %s", app)
			fmt.Fprintf(w, "%s will run immediately\n", app)
			a.runRealTime(app)
			return
		}
		a.addPending(app)
		log.Println(a.apps)
		fmt.Fprintf(w, "ok, %s will run in 00:00:00\n", app)
	})
	go func() {
		log.Println("server listen on 8080")
		error := http.ListenAndServe(":8080", nil)
		if error != nil {
			log.Fatal(error)
		}
	}()
}

func (a *AppDelayRunner) dump() {
	log.Printf("ready to dump")
	if len(a.apps) == 0 {
		log.Printf("no app exist")
		ioutil.WriteFile("apps.conf", []byte(""), 0644)
	} else {
		res := make([]string, 0)
		for v := range a.apps {
			res = append(res, v)
		}
		apps := strings.Join(res, ",")
		log.Printf("save apps: %s", apps)
		ioutil.WriteFile("apps.conf", []byte(apps), 0644)
	}
}

func (a *AppDelayRunner) waitStop() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-c:
			a.dump()
			log.Fatal("exit")
		}
	}()
	// a.wg.Wait()
}
func (a *AppDelayRunner) readFile() {
	content, err := ioutil.ReadFile("apps.conf")
	if err != nil {
		log.Println("apps.conf not found")
		return
	}
	s := string(content[:])
	if s == "" {
		return
	}
	log.Printf("load apps: %s", s)
	for _, app := range strings.Split(string(content[:]), ",") {
		a.addPending(app)
	}
}

// Start listen and run app deploy
func (a *AppDelayRunner) Start() {
	a.readFile()
	a.startServer()
	a.waitStop()
}

func CheckAvailable(address string) {
	tc := time.NewTicker(15 * time.Second)
	done := make(chan bool, 1)
	for {
		select {
		case <-tc.C:
			resp, error := http.Get(address)
			if error != nil {
				log.Printf("deploy server not ok, %v", error)
				tc.Stop()
				done <- true
			} else {
				resp.Body.Close()
				log.Println("deploy server check ok")
			}
		case <-done:
			close(done)
			startAsService()
			return
		}
	}
}

func startAsService() {
	New().Start()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var p = flag.String("address", "http://m1.epj:8080", "deploy server address")
	CheckAvailable(*p)
	wg.Wait()
}
