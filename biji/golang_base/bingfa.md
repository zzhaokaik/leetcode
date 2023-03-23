#GOland 并发模型

###线程间通信方式
线程间存在两种通讯方式  
1.共享内存  
2.消息传递  
同时通过内存获取数据

###共享内存
抽象层级：抽象层级低。对资源进行细粒度或对性能有高要求使用
耦合：高
线程竞争：需要加锁

###发送消息
抽象层级：高，提供封装和耦合设计（channel）
耦合：低
线程竞争：同一时间只有1个线程访问

goland 2种并发模型
1.共享内存模型
2.csp模型

###共享内存模型
![img.png](img.png)

###CSP并发模型
通过发送消息的方式来同步信息，Go语言推荐使用的通信顺序进程（communicating sequential processes）并发模型，通过goroutine和channel来实现

goroutine 是Go语言中并发的执行单位，可以理解为”线程“
channel是Go语言中各个并发结构体(goroutine)之前的通信机制。 通俗的讲，就是各个goroutine之间通信的”管道“，类似于Linux中的管道
![img_1.png](img_1.png)
##Go 有哪些并发同步原语
###原子操作
Mutex、RWMutex 等并发原语的底层实现是通过 atomic 包中的一些原子操作来实现的，原子操作是最基础的并发原语
![img_2.png](img_2.png)
``` 
package main

import (
    "fmt"
    "sync/atomic"
)

var opts int64 = 0

func main() {
    add(&opts, 3)
    load(&opts)
    compareAndSwap(&opts, 3, 4)
    swap(&opts, 5)
    store(&opts, 6)
}

func add(addr *int64, delta int64) {
    atomic.AddInt64(addr, delta) //加操作
    fmt.Println("add opts: ", *addr)
}

func load(addr *int64) {
    fmt.Println("load opts: ", atomic.LoadInt64(&opts))
}

func compareAndSwap(addr *int64, oldValue int64, newValue int64) {
    if atomic.CompareAndSwapInt64(addr, oldValue, newValue) {
        fmt.Println("cas opts: ", *addr)
        return
    }
}

func swap(addr *int64, newValue int64) {
    atomic.SwapInt64(addr, newValue)
    fmt.Println("swap opts: ", *addr)
}

func store(addr *int64, newValue int64) {
    atomic.StoreInt64(addr, newValue)
    fmt.Println("store opts: ", *addr)
}

```

###Channel
channel 管道，高级同步原语，goroutine之间通信的桥梁

使用场景：消息队列、数据传递、信号通知、任务编排、锁

``` 
package main

import (
    "fmt"
    "time"
)

func main() {
    c := make(chan struct{}, 1)
    for i := 0; i < 10; i++ {
        go func() {
            c <- struct{}{}
            time.Sleep(1 * time.Second)
            fmt.Println("通过ch访问临界区")
            <-c
        }()
    }
    for {
    }
}

```

###基本并发原语
Go 语言在 sync包中提供了用于同步的一些基本原语，这些基本原语提供了较为基础的同步功能，但是它们是一种相对原始的同步机制，在多数情况下，我们都应该使用抽象层级更高的 Channel 实现同步。

常见的并发原语如下：sync.Mutex、sync.RWMutex、sync.WaitGroup、sync.Cond、sync.Once、sync.Pool、sync.Context

####sync.Mutex
sync.Mutex （互斥锁） 可以限制对临界资源的访问，保证只有一个 goroutine 访问共享资源

使用场景：大量读写，比如多个 goroutine 并发更新同一个资源，像计数器
``` 
package main

import (
    "fmt"
    "sync"
)

func main() {
    // 封装好的计数器
    var counter Counter
    var wg sync.WaitGroup
    var gNum = 1000
    wg.Add(gNum)
    // 启动10个goroutine
    for i := 0; i < gNum; i++ {
        go func() {
            defer wg.Done()
            counter.Incr() // 受到锁保护的方法
        }()
    }
    wg.Wait()
    fmt.Println(counter.Count())
}

// 线程安全的计数器类型
type Counter struct {
    mu    sync.Mutex
    count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) Count() uint64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

####sync.RWMutex
sync.RWMutex （读写锁） 可以限制对临界资源的访问，保证只有一个 goroutine 写共享资源，可以有多个goroutine 读共享资源

使用场景：大量并发读，少量并发写，有强烈的性能要求

``` 
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // 封装好的计数器
    var counter Counter
    var gNum = 1000
    // 启动10个goroutine
    for i := 0; i < gNum; i++ {
        go func() {
            counter.Count() // 受到锁保护的方法
        }()
    }
    for { // 一个writer
        counter.Incr() // 计数器写操作
        fmt.Println("incr")
        time.Sleep(time.Second)
    }
}

// 线程安全的计数器类型
type Counter struct {
    mu    sync.RWMutex
    count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) Count() uint64 {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.count
}

```

####sync.WaitGroup
sync.WaitGroup 可以等待一组 Goroutine 的返回

使用场景：并发等待，任务编排，一个比较常见的使用场景是批量发出 RPC 或者 HTTP 请求

``` 
requests := []*Request{...}
wg := &sync.WaitGroup{}
wg.Add(len(requests))

for _, request := range requests {
    go func(r *Request) {
        defer wg.Done()
        // res, err := service.call(r)
    }(request)
}
wg.Wait()

```

####sync.Cond
sync.Cond 可以让一组的 Goroutine 都在满足特定条件时被唤醒

使用场景：利用等待 / 通知机制实现阻塞或者唤醒
``` 
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

var status int64

func main() {
    c := sync.NewCond(&sync.Mutex{})
    for i := 0; i < 10; i++ {
        go listen(c)
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

func listen(c *sync.Cond) {
    c.L.Lock()
    for atomic.LoadInt64(&status) != 1 {
        c.Wait()
    }
    fmt.Println("listen")
    c.L.Unlock()
}
```

####sync.Once
sync.Once 可以保证在 Go 程序运行期间的某段代码只会执行一次

使用场景：常常用于单例对象的初始化场景
``` 
package main

import (
    "fmt"
    "sync"
)

func main() {
    o := &sync.Once{}
    for i := 0; i < 10; i++ {
        o.Do(func() {
            fmt.Println("only once")
        })
    }
}
```

####sync.Pool
sync.Pool可以将暂时将不用的对象缓存起来，待下次需要的时候直接使用，不用再次经过内存分配，复用对象的内存，减轻 GC 的压力，提升系统的性能（频繁地分配、回收内存会给 GC 带来一定的负担，严重的时候会引起 CPU 的毛刺）

使用场景：对象池化， TCP连接池、数据库连接池、Worker Pool
``` 
package main

import (
    "fmt"
    "sync"
)

func main() {
    pool := sync.Pool{
        New: func() interface{} {
            return 0
        },
    }

    for i := 0; i < 10; i++ {
        v := pool.Get().(int)
        fmt.Println(v) // 取出来的值是put进去的，对象复用；如果是新建对象，则取出来的值为0
        pool.Put(i)
    }
}

```

####sync.Map
sync.Map 线程安全的map

使用场景：map 并发读写

``` 
package main

import (
    "fmt"
    "sync"
)

func main() {
    var scene sync.Map
    // 将键值对保存到sync.Map
    scene.Store("1", 1)
    scene.Store("2", 2)
    scene.Store("3", 3)
    // 从sync.Map中根据键取值
    fmt.Println(scene.Load("1"))
    // 根据键删除对应的键值对
    scene.Delete("1")
    // 遍历所有sync.Map中的键值对
    scene.Range(func(k, v interface{}) bool {
        fmt.Println("iterate:", k, v)
        return true
    })
}

```

####sync.Context
sync.Context 可以进行上下文信息传递、提供超时和取消机制、控制子 goroutine 的执行

使用场景：取消一个goroutine的执行
``` 
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        defer func() {
            fmt.Println("goroutine exit")
        }()
        for {
            select {
            case <-ctx.Done():
                fmt.Println("receive cancel signal!")
                return
            default:
                fmt.Println("default")
                time.Sleep(time.Second)
            }
        }
    }()
    time.Sleep(time.Second)
    cancel()
    time.Sleep(2 * time.Second)
}

```