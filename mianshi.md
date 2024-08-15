### 1 三次握手是TCP协议中建立连接的过程，确保客户端和服务器之间的连接是可靠的。以下是三次握手的详细步骤：

第一次握手（SYN）：

客户端向服务器发送一个SYN（同步序列编号）包，表示客户端希望建立连接，并且包含一个初始序列号。
第二次握手（SYN-ACK）：

服务器收到SYN包后，向客户端发送一个SYN-ACK包，表示同意建立连接，并且包含服务器的初始序列号和对客户端SYN包的确认序列号。
第三次握手（ACK）：

客户端收到SYN-ACK包后，向服务器发送一个ACK（确认）包，表示确认收到服务器的SYN-ACK包，并且包含对服务器SYN包的确认序列号。
完成三次握手后，客户端和服务器之间的连接就建立起来了，可以开始数据传输

### http1 和http2的区别？
HTTP/1.1 和 HTTP/2 是两种不同版本的超文本传输协议（HTTP），它们在性能和功能上有显著的区别。以下是一些主要区别：

多路复用（Multiplexing）：

HTTP/1.1：每个请求/响应对需要一个单独的TCP连接，容易导致“队头阻塞”（Head-of-line blocking）。
HTTP/2：引入了多路复用技术，允许在一个TCP连接上并行发送多个请求和响应，减少了延迟和资源消耗。
头部压缩（Header Compression）：

HTTP/1.1：HTTP头部信息未压缩，导致冗余数据传输。
HTTP/2：使用HPACK压缩算法对头部信息进行压缩，减少了传输的数据量，提高了传输效率。
服务器推送（Server Push）：

HTTP/1.1：服务器只能响应客户端的请求，不能主动推送资源。
HTTP/2：服务器可以主动向客户端推送资源，减少了客户端请求的延迟。
数据帧（Data Frames）：

HTTP/1.1：数据以纯文本形式传输。
HTTP/2：数据以二进制帧的形式传输，更高效且更易于解析。
连接管理（Connection Management）：

HTTP/1.1：需要频繁建立和关闭TCP连接，增加了开销。
HTTP/2：一个TCP连接可以处理多个请求和响应，减少了连接的开销。
优先级和流量控制（Priority and Flow Control）：

HTTP/1.1：没有内置的优先级和流量控制机制。
HTTP/2：支持请求的优先级和流量控制，允许客户端和服务器更好地管理资源。
### 协程池
package main

import (
    "fmt"
    "sync"
)

// Task 代表一个任务
type Task struct {
    ID      int
    Execute func() error
}

// GoroutinePool 代表一个协程池
type GoroutinePool struct {
    taskQueue   chan Task
    workerCount int
    wg          sync.WaitGroup
}

// NewGoroutinePool 创建一个新的协程池
func NewGoroutinePool(workerCount int, taskQueueSize int) *GoroutinePool {
    return &GoroutinePool{
        taskQueue:   make(chan Task, taskQueueSize),
        workerCount: workerCount,
    }
}

// Start 启动协程池
func (p *GoroutinePool) Start() {
    for i := 0; i < p.workerCount; i++ {
        p.wg.Add(1)
        go p.worker(i)
    }
}

// worker 是协程池中的工作协程
func (p *GoroutinePool) worker(id int) {
    defer p.wg.Done()
    for task := range p.taskQueue {
        fmt.Printf("Worker %d processing task %d\n", id, task.ID)
        if err := task.Execute(); err != nil {
            fmt.Printf("Task %d failed: %v\n", task.ID, err)
        }
    }
}

// Submit 提交任务到协程池
func (p *GoroutinePool) Submit(task Task) {
    p.taskQueue <- task
}

// Stop 停止协程池
func (p *GoroutinePool) Stop() {
    close(p.taskQueue)
    p.wg.Wait()
}

func main() {
    // 创建一个协程池，包含3个工作协程，任务队列大小为10
    pool := NewGoroutinePool(3, 10)

    // 启动协程池
    pool.Start()

    // 提交任务到协程池
    for i := 0; i < 10; i++ {
        taskID := i
        pool.Submit(Task{
            ID: taskID,
            Execute: func() error {
                fmt.Printf("Executing task %d\n", taskID)
                return nil
            },
        })
    }

    // 停止协程池
    pool.Stop()
}
### tcp 和udp的区别
1. 连接方式
TCP：面向连接的协议，在传输数据之前需要建立连接（三次握手），传输结束后需要断开连接（四次挥手）。
UDP：无连接的协议，发送数据之前不需要建立连接，直接发送数据。
2. 可靠性
TCP：提供可靠的数据传输，保证数据包按顺序到达且不丢失。通过重传机制和确认机制确保数据的完整性。
UDP：不保证数据包的可靠性，数据包可能会丢失、重复或乱序到达。没有重传机制和确认机制。
3. 流量控制和拥塞控制
TCP：具有流量控制和拥塞控制机制，能够根据网络状况调整数据传输速率，避免网络拥塞。
UDP：没有流量控制和拥塞控制机制，发送数据的速率不受网络状况影响。
4. 传输效率
TCP：由于需要建立连接、确认数据包和进行流量控制，传输效率相对较低。
UDP：由于不需要建立连接和确认数据包，传输效率较高，适用于实时性要求高的应用。
5. 应用场景
TCP：适用于对数据传输可靠性要求高的应用，如网页浏览（HTTP/HTTPS）、文件传输（FTP）、电子邮件（SMTP/POP3）等。
UDP：适用于对实时性要求高且可以容忍少量数据丢失的应用，如视频直播、在线游戏、语音通话（VoIP）等。
### grpc 和json的区别
1. 数据格式
gRPC：
使用 Protocol Buffers（protobuf）作为序列化格式，数据以二进制形式传输。
优点：高效、紧凑，适合高性能需求的场景。
JSON：
使用文本格式进行数据交换，数据以人类可读的字符串形式传输。
优点：易于阅读和调试，广泛支持。
2. 通信协议
gRPC：
基于 HTTP/2 协议，支持多路复用、流量控制、头部压缩和服务器推送等特性。
优点：高效的网络传输，适合实时通信和高并发场景。
JSON：
通常基于 HTTP/1.1 协议进行传输。
优点：简单易用，广泛支持。
3. 接口定义
gRPC：
使用 .proto 文件定义服务和消息结构，支持自动生成客户端和服务器代码。
优点：强类型定义，减少错误，提升开发效率。
JSON：
没有统一的接口定义方式，通常使用 RESTful API 设计规范。
优点：灵活性高，但需要手动编写接口文档和代码。
4. 性能
gRPC：
由于使用二进制格式和 HTTP/2 协议，性能较高，延迟较低。
适合高性能和低延迟需求的应用，如微服务通信、实时系统等。
JSON：
由于使用文本格式和 HTTP/1.1 协议，性能相对较低，延迟较高。
适合对性能要求不高的应用，如Web应用、移动应用等。
5. 应用场景
gRPC：
适用于微服务架构、实时通信、高性能系统等场景。
例如：内部服务间通信、实时数据流处理等。
JSON：
适用于Web应用、移动应用、公开API等场景。
例如：前后端数据交换、第三方API集成等
### golang的垃圾回收
1. 垃圾回收器的特点
并发：Go的垃圾回收器是并发的，能够在程序运行时进行垃圾回收操作，减少对程序性能的影响。
三色标记法：Go垃圾回收器使用三色标记法（Tri-color Mark and Sweep）来标记和清除不再使用的对象。
分代回收：Go垃圾回收器采用分代回收策略，将对象分为不同的代，根据对象的生命周期优化回收策略。
2. 垃圾回收的工作原理
标记阶段：垃圾回收器从根对象（全局变量、栈变量等）开始，递归遍历所有可达对象，并将其标记为“活跃”。
清除阶段：垃圾回收器遍历堆内存，清除所有未被标记为“活跃”的对象，释放内存。
3. 调优和监控
GOGC 环境变量：可以通过设置 GOGC 环境变量来调整垃圾回收的频率。默认值为100，表示当堆内存增长100%时触发垃圾回收。
运行时监控：Go提供了 runtime 包，可以用于监控和调优垃圾回收。例如，使用 runtime.ReadMemStats 获取内存统计信息。
### redis 消息队列
使用 List 实现消息队列
LPUSH + BRPOP：生产者使用 LPUSH 命令将消息推入队列，消费者使用 BRPOP 命令阻塞式地弹出消息。
优点：简单易用，适合大多数场景。
缺点：不支持消息确认和重试机制。
使用 Pub/Sub 实现发布/订阅模式
PUBLISH + SUBSCRIBE：生产者使用 PUBLISH 命令发布消息，消费者使用 SUBSCRIBE 命令订阅消息。
优点：支持一对多的消息传递，适合广播消息的场景。
缺点：消息不持久化，消费者离线时会丢失消息。
使用 Stream 实现高级消息队列
XADD + XREAD：生产者使用 XADD 命令添加消息到流，消费者使用 XREAD 命令读取消息。
优点：支持消息持久化、消息确认、消费组等高级特性。
缺点：相对复杂，适合对消息队列有高级需求的场景。

### 高并发扣减库存
数据库事务：适合简单场景，但在高并发下可能会有性能瓶颈。
分布式锁：适合需要严格控制并发的场景，但需要注意锁的粒度和性能。
消息队列：适合高并发和异步处理的场景，但需要处理消息的顺序和幂等性。