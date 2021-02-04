package concurrent

import (
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

//如何通过通道来停止goroutine
func TestStopGoroutineByChan(t *testing.T) {

}

//如何通过Wait Group来停止goroutine
func TestStopGoroutineByWaitGroup(t *testing.T) {

}

//如何设置goroutine的最大cpu数量
func TestUseGoMaxProcess(t *testing.T) {

}

//如何使用Goroutine的局部存储
func TestCreateGoroutineTLS(t *testing.T) {

}

//使用Gosched释放线程先去执行别的任务
func TestGoRoutineByGosched(t *testing.T) {

}

//使用Goexit来终止当前任务

func TestGoRoutineByGoexit(t *testing.T) {

}
