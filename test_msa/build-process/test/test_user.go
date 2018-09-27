package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Event struct {
	NameID    string
	eventType string
	sign      string
	info      string
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(3)

	// 标记的A事件cache
	ACache := make(map[string]Event)
	// B事件cache
	BCache := make(map[string]Event)

	eventChan := make(chan Event, 100)
	BChan := make(chan Event, 10)

	go func() {
		for {
			select {
			case e := <-eventChan:
				log.Println("eventChan:进入事件事件处理")
				if e.eventType == "A" {
					//处理A
					e.sign = "done"
					e.info = "A finished"
					log.Printf("eventChan: 处理A事件，nameId: %s \n", e.NameID)
					if bevent, ok := BCache[e.NameID]; ok {
						log.Printf("eventChan: nameId: %s ,从BCache中找到B事件，抛入Bchan\n", e.NameID)
						BChan <- bevent
					} else {
						ACache[e.NameID] = e
						log.Printf("eventChan: nameId: %s ，BCache中没有找到B事件，将自己抛入标记ACache\n", e.NameID)
						time.AfterFunc(1*time.Hour, func() {
							delete(ACache, e.NameID)
						})
					}
				} else if e.eventType == "B" {
					log.Println("eventChan: 事件类型为B")
					if _, ok := ACache[e.NameID]; ok {
						log.Printf("eventChan: nameId: %s ,ACache中找到,A事件执行过了,抛入Bchan\n", e.NameID)
						BChan <- e
					} else {
						log.Printf("eventChan: nameId: %s ,ACache中没有找到,A事件没有执行,抛入BCache\n", e.NameID)
						BCache[e.NameID] = e
					}
				}
			case b := <-BChan:
				log.Printf("进入B事件, nameId: %s", b.NameID)
				go func() {
					b.info = "B finished"
					delete(BCache, b.NameID)
					delete(ACache, b.NameID)
				}()
			}
		}
	}()

	go func() {
		helloA := Event{
			"hello",
			"A",
			"no",
			"no",
		}
		eventChan <- helloA
		for len := 20; len > 0; len-- {
			eventChan <- Event{
				fmt.Sprintf("e-%v", len),
				"A",
				"no",
				"no",
			}
		}

	}()

	go func() {
		helloB := Event{
			"hello",
			"B",
			"*",
			"no",
		}
		eventChan <- helloB
		for len := 20; len > 0; len-- {
			eventChan <- Event{
				fmt.Sprintf("e-%v", len),
				"B",
				"no",
				"no",
			}
		}
	}()

	wg.Wait()
}
