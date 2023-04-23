package base

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//var wg sync.WaitGroup
//
//func main() {
//	ch1 := make(chan struct{}, 1)
//	ch2 := make(chan struct{}, 1)
//	ch3 := make(chan struct{}, 1)
//	ch1 <- struct{}{}
//	wg.Add(3)
//	start := time.Now().Unix()
//	go print("gorouine1", ch1, ch2)
//	go print("gorouine2", ch2, ch3)
//	go print("gorouine3", ch3, ch1)
//	wg.Wait()
//	end := time.Now().Unix()
//	fmt.Printf("duration:%d\n", end-start)
//}
//
//func print(gorouine string, inputchan chan struct{}, outchan chan struct{}) {
//	// 模拟内部操作耗时
//	time.Sleep(1 * time.Second)
//	select {
//	case <-inputchan:
//		fmt.Printf("%s\n", gorouine)
//		outchan <- struct{}{}
//	}
//	wg.Done()
//}


var status int64

func main1() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c,i)
	}
	time.Sleep(1 * time.Second)
	go broadcast(c)
	time.Sleep(1 * time.Second)
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Signal()
	c.L.Unlock()
}

func listen(c *sync.Cond,a int) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Println("listen",a)
	c.L.Unlock()
}

