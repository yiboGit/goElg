package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-ticker.C:
			log.Printf("日志记录测试--%v \n", time.Now().Format("2006-01-02 : 15:04:05"))
		}
	}
}
