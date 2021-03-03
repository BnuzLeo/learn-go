package concurrent

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

//什么是goroutine
func TestWhatIsGoroutine(t *testing.T) {
	/*
		并发：逻辑上具备同时处理多个任务的能力
		并行：物理上在同一时刻执行多个任务的能力
	*/

	/*
		知识点：
		1. go 只是创建了一个任务单元进入系统队列进行运行
		2. go routine默认的栈内存空间只有2kb，而线程栈的默认空间是mb级别，但GoRoutine在需要情况下，最大可以是GB
	*/

	go println("Hi, I am go routine")

	go func() { println("Hi, I am go routine in function") }()

	time.Sleep(time.Second * 1)
}

//goroutine的延迟执行, 立即计算并且计算执行参数
func TestGoRoutineDelay(t *testing.T) {
	a := 100

	go func(x int, y int) {
		fmt.Println("go: ", x, y) // 100, 1
		time.Sleep(time.Second * 1)
	}(a, Counter()) // 这个时候a=100，然后马上计算出Counter()是1

	a += 100
	fmt.Println("main:", a, Counter()) // 200, 2
}

//进程退出时不会等待并发任务结束，可以通过channel来进行阻塞，然后发出退出信号
func TestStopGoroutineByChan(t *testing.T) {
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("Go Run")
		close(exit)
	}()

	fmt.Println("main ...")
	<-exit
	fmt.Println("main exit...")
}

//Wait Group: 通过wait group来进行阻塞
func TestStopGoroutineByWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println("goroutine", id, "done")
		}(i)
	}

	fmt.Println("main enter")
	wg.Wait()
	fmt.Println("main exit")
}

//GOMAXPROCS :如何设置goroutine的最大cpu数量
func TestUseGoMaxProcess(t *testing.T) {
	_ = runtime.GOMAXPROCS(1)
	// 设置完后，即便你有4核，也只会用一核来运算

	// 4核运算
	// real : 0.3s 程序执行时间
	// total： 1.2s 多核执行累计时间

	// 1核运算
	// real : 1.2s 程序执行时间
	// total： 1.2s 多核执行累计时间
}

//Local Storage: 如何使用Goroutine的局部存储
/*
	TLS : Thread local storage(线程局部存储)
	go routine没有优先级，甚至连返回值都被抛弃。
	但是除了优先级以外，其他都可以实现
*/
func TestCreateGoroutineTLS(t *testing.T) {
	var wg sync.WaitGroup
	var gs [5]struct {
		id     int
		result int
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i)
	}

	wg.Wait()
	fmt.Printf("%v\n", gs)
}

//GoSched： 暂停，释放线程去执行其他任务。当前任务被放回队列，等下次执行的时候再次执行
func TestGoRoutineByGosched(t *testing.T) {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func() {
			fmt.Println("func 1...")
		}()

		for i := 0; i < 4; i++ {
			fmt.Println("func 2 :", i)
			if i == 1 {
				runtime.Gosched()
			}
		}
	}()
	<-exit
}

//GoExit：终止任务
/*
	如果在方法里面执行：他会马上终止当前并发任务，不会影响延迟执行

	如果在main方法里面执行： 他会等待其他任务完成，然后结束进程
*/
func TestGoRoutineByGoexit(t *testing.T) {
	exit := make(chan struct{})

	go func() {
		defer close(exit)           // 执行
		defer fmt.Println("func 1") // 执行

		func() {
			defer func() {
				fmt.Println("func 2", recover() == nil)
			}() // 执行

			func() {
				fmt.Println("func 3") // 执行
				runtime.Goexit()
				fmt.Println("func 3 done.") // 不执行
			}()

			fmt.Println("func 2 done.") // 不执行
		}()
		fmt.Println("func1 done.") // 不执行
	}()

	<-exit
}
