package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

// StaffMsgResponse response of staff message
type StaffMsgResponse struct {
	ID      int64
	ErrType string
	Error   string
}

// serveWs handles websocket requests from the peer.
func serveWs(ex *Exchanger, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{ex: ex, conn: conn, ID: 0, send: make(chan []byte, 1024)}
	client.ex.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func main() {
	flag.Parse()
	ex := newExchanger()
	go ex.run()
	// websocket
	http.HandleFunc("/im/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(ex, w, r)
	})
	// 接入
	http.HandleFunc("/im/user/bind", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data, error := ioutil.ReadAll(r.Body)
		if error != nil {
			log.Printf("body read error: %v", error)
			http.Error(w, "", 500)
			return
		}
		log.Println(string(data))
		ex.userBind <- data
	})
	// 发送消息给公众号用户
	http.HandleFunc("/im/user/staff-message", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data, error := ioutil.ReadAll(r.Body)
		if error != nil {
			log.Printf("body read error: %v", error)
			json.NewEncoder(w).Encode(StaffMsgResponse{ID: 0, ErrType: "request", Error: error.Error()})
			return
		}
		id, err := ex.handleStaffMessage(data)
		res := StaffMsgResponse{ID: id, ErrType: err.errType}
		if err.err != nil {
			res.Error = err.err.Error()
		}
		json.NewEncoder(w).Encode(res)
	})
	// 公众号用户发来消息
	http.HandleFunc("/im/user", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data, error := ioutil.ReadAll(r.Body)
		if error != nil {
			log.Printf("body read error: %v", error)
			http.Error(w, "", 500)
			return
		}
		log.Println(string(data))
		ex.userMessages <- data
	})

	log.Printf("im server start")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
