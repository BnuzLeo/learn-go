package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
	问题：Go并没有严格的并发安全，全局变量，指针，引用类型这些非安全内存可以共享操作
	解决方案： 开发人员自行维护数据一致，鼓励使用CSP，通信代替共享内存，实现并发安全
	CSP（Communication Sequential Process）
*/

/*
	通道阻塞：同步模式必须要配对的go routine出现，否则会一致阻塞

	而异步模式在缓冲区未满或者数据未读完前不会阻塞
*/

func TestBlock(t *testing.T) {
	ints := make(chan int, 3) // 弄3个通道出来

	ints <- 1 // 未满不会阻塞
	ints <- 2

	fmt.Println(<-ints) // 未读完不会阻塞
	fmt.Println(<-ints)
}

/*
	收发：有三种，普通模式，ok-idom，range
*/

// 普通模式
func TestCommunication(t *testing.T) {
	c := make(chan string)

	go func() {
		defer close(c)
		c <- "hi"
		fmt.Println("Say hi to c")
	}()

	receive := <-c
	fmt.Println("receive: ", receive)
}

// ok-idom
func TestOkIDom(t *testing.T) {
	done := make(chan struct{})
	ints := make(chan int)

	go func() {
		defer close(done)

		for {
			x, ok := <-ints
			if !ok {
				return
			}
			println(x)
		}
	}()

	ints <- 1
	ints <- 2
	ints <- 3
	close(ints) // 及时用close，不然会"all goroutines are asleep - deadlock!"

	<-done
}

// range
func TestRange(t *testing.T) {
	done := make(chan struct{})
	ints := make(chan int)

	go func() {
		defer close(done)

		for i := range ints {
			fmt.Println(i)
		}
	}()

	ints <- 1
	ints <- 2
	ints <- 3
	close(ints)

	<-done
}

/*
	广播: 通知可以是群体性的
*/

func TestBroadcast(t *testing.T) {
	var wg sync.WaitGroup
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println(id, "Ready")
			<-done
			fmt.Println(id, "Run")
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Ready? GO!")
	close(done)
	wg.Wait()
}

/*
	单向通道

	send : chan<- int
	recv : <-chan int

	不可以做的事情：
	1. 不可以做逆向的事情，例如从send中获取数据 <-send
	2. 不可以关闭接收端， close(recv)
	3. 不可以转化回双向通道， c := (chan int)(recv)
*/

func TestOnWay(t *testing.T) {
	var wg sync.WaitGroup

	c := make(chan int)
	var send chan<- int = c // send 是 c通道的发送方
	var recv <-chan int = c // recv 是 c通道的接收方

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(c)

		send <- 1
		send <- 2
		send <- 3
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for r := range recv {
			fmt.Println(r)
		}
	}()

	wg.Wait()
}
