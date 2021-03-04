package concurrent

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var num int

func Counter() int {
	num++
	return num
}

// 模式
type Receiver struct {
	sync.WaitGroup
	data chan int
}

func NewReceiver() *Receiver {
	r := &Receiver{
		data: make(chan int),
	}
	r.Add(1)

	go func() {
		defer r.Done()
		for x := range r.data {
			fmt.Println("recv： ", x)
		}
	}()

	return r
}

// 通道池
type pool chan []byte

func GetPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p: // 返回成功
	default:
		v = make([]byte, 10) // 返回失败，新建
	}
	return v
}

func (p pool) put(b []byte) bool {
	select {
	case p <- b: // 放回
	default:
		return false // 放回失败，return false
	}
	return true
}

// 退出处理器
type Exit struct {
	sync.RWMutex
	fs      []func()
	signals chan os.Signal
}

func (e *Exit) AtExit(f func()) {
	e.Lock()
	defer e.Unlock()
	e.fs = append(e.fs, f)
}

func (e *Exit) WaitExit() {

	if e.signals == nil {
		e.signals = make(chan os.Signal)
		signal.Notify(e.signals, syscall.SIGINT, syscall.SIGTERM)
	}

	e.RLock()
	for _, f := range e.fs {
		defer f()
	}
	e.RUnlock()
}
