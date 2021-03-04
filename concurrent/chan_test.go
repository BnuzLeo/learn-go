package concurrent

import (
	"fmt"
	"os"
	"runtime"
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

/*
	Select: 如果同时处理多个通道，可以选用select语句。 它会随机找一个通道进行收发
*/
func TestSelectForRandom(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()
		var (
			name string
			x    int
			ok   bool
		)

		for {
			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			fmt.Println(name, x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 1; i < 5; i++ {
			select {
			case a <- i:
			case a <- i * 10:
			}
		}
	}()

	wg.Wait()
}

/*
	Select: 如果想等所有的通道消息处理完，可以把通道设置为nil，这样通道会被阻塞
*/

func TestSelectBlock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()

		for {
			select {
			case x, ok := <-a:
				if !ok {
					a = nil
					break
				}
				name := "a"
				fmt.Println(name, x)
			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}
				name := "b"
				fmt.Println(name, x)
			}
			if a == nil && b == nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)

		for i := 1; i < 3; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()

		defer close(b)

		for i := 1; i < 3; i++ {
			b <- i * 10
		}
	}()

	wg.Wait()
}

/*
	Select: 所有通道都不可行的时候，可以通过default来避开select的阻塞
*/

func TestSelectUnBlock(t *testing.T) {
	done := make(chan struct{})
	num := make(chan int)

	go func() {
		defer close(done)
		for {
			select {
			case x, ok := <-num:
				if !ok {
					return
				}
				fmt.Println("Date :", x)
			default:

			}
			fmt.Println("Date : ", time.Now())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	num <- 100
	num <- 200
	close(num)

	<-done
}

/*
	模式： 通过工厂方法将goroutine和通道绑定
*/

func TestMode(t *testing.T) {
	r := NewReceiver()
	r.data <- 1
	r.data <- 2
	r.data <- 3
	close(r.data)
	r.Wait()
}

/*
	通道池： 通道本来就是一个并发安全的队列，可以用作Id generator， pool等用途
*/
func TestGetPool(t *testing.T) {
	pool := GetPool(3)
	isSuccess := pool.put([]byte{'a'})
	fmt.Println("a", isSuccess)
	isSuccess = pool.put([]byte{'b'})
	fmt.Println("b", isSuccess)
	isSuccess = pool.put([]byte{'c'})
	fmt.Println("c", isSuccess)
	isSuccess = pool.put([]byte{'d'})
	fmt.Println("d", isSuccess) // 因为通道同时最多是3个同时在线
	fmt.Println(pool.get())
	fmt.Println(pool.get())
	fmt.Println(pool.get())
	fmt.Println(pool.get()) //[0 0 0 0 0 0 0 0 0 0]
	close(pool)
}

/*
	信号量：用通道实现信号量
*/

func TestSemaphore(t *testing.T) {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	sem := make(chan int, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()
			sem <- x                 // 添加信号
			time.Sleep(time.Second)  // 做你要做的事情
			defer func() { <-sem }() // 释放信号
			fmt.Println(x, time.Now())
		}(i)
	}

	wg.Wait()
}

/*
	标准库Time的 timeout 和 tick 实现
*/

func TestTimeChannel(t *testing.T) {
	// Time out channel
	go func() {
		for {
			select {
			case <-time.After(time.Second * 5): // 5秒之后，毁天灭地
				fmt.Println("timeout")
				os.Exit(0)
			}
		}
	}()

	// Time ticket channel
	go func() {
		for {
			select {
			case <-time.Tick(time.Second): // 每过1s，执行一次
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil) // 空通道直接阻塞
}

/*
	退出处理器：atexit
*/

func TestAtExit(t *testing.T) {
	exit := &Exit{}
	exit.AtExit(func() {
		fmt.Println("exit 1....")
	})
	exit.AtExit(func() {
		fmt.Println("exit 2....")
	})

	exit.WaitExit()

}
