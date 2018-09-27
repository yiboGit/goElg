package main
import (
	"sync"
	"fmt"
	// "sync/atomic"
	"time"
)

type Working struct {
	l sync.Mutex
	lastTime int64
	data int
}

func (w *Working) Add(data int)bool {
	now := time.Now().Unix()
	if now - w.lastTime < 1000 {
		fmt.Print("<1\n")
		return false
	}
	w.l.Lock()
	now = time.Now().Unix()
	fmt.Printf("get lock %d now %d, %d \n", data, now, w.lastTime)
	if now - w.lastTime < 1 {
		fmt.Printf("lock <1 %d \n", data)
		w.l.Unlock()
		return false
	}
	w.lastTime = now
	w.data = data
	fmt.Print("set ok", data)
	w.l.Unlock()
	return true
}


func main1() {
	var s sync.WaitGroup
	num := 20
	s.Add(num)
	w := Working{lastTime: time.Now().UnixNano()}
	for i:= 1; i<=num; i++ {
		go func(t int) {
			defer s.Done()
			d := time.Millisecond * 100 * time.Duration(t)
			time.Sleep(d)
			w.Add(t)
		}(i)
	}
	s.Wait()
	fmt.Print(w.data)
}