package method

import (
	"fmt"
	"sync"
	"testing"
)

/*
	Go的函数是一等公民：
	Go的函数是第一对象，可以用作函数的入参和返回值，也可以存入对象实体
*/

func format(str string) (format string) {
	result := str + " format"
	fmt.Println(result)
	return result
}

func TestFirstObject(t *testing.T) {
	Format(format, "first object")
}

/*
	不管是指针，引用类型还是其他类型参数，其实都是"引用传递"
	区别无非是拷贝目标对象还是拷贝指针而已。

	拷贝目标和拷贝指针哪个比较好？
	答：其实要分情况。如果拷贝的目标比较小或者在并发情况下，拷贝目标会比较好。因为拷贝指针有可能会把对象放入堆内存，成本要把内存分配和内存回收计算进去
*/
func f(x, y int, s string, _ bool) (*int, error) {
	z := x + y
	fmt.Println(z, s)
	return &z, nil
}

func TestParameterAndReturn(t *testing.T) {
	f(1, 2, "3", true)
}

/*
	闭包：上下文中引用了自由变量的函数。他是函数和环境变量的一个集合
*/

func closure1(i int) func() {
	return func() { fmt.Println(i) }
}

func closure2() []func() {
	var fs []func()
	for i := 0; i < 2; i++ {
		fs = append(fs, func() {
			fmt.Println(&i, i)
		})
	}
	return fs
}

func closure3() []func() {
	var fs []func()
	for i := 0; i < 2; i++ {
		x := i
		fs = append(fs, func() {
			fmt.Println(&x, x)
		})
	}
	return fs
}

func closure4() []func(x int) {
	var fs []func(x int)
	for i := 0; i < 2; i++ {
		func(x int) {
			fmt.Println(&x, x)
		}(i)
	}
	return fs
}

func TestClosure(t *testing.T) {
	/*
		普通闭包：返回的函数引用了环境变量i，所以他是一个闭包。
		闭包带来的问题：就是会把数据引用到堆内存
	*/
	f1 := closure1(1)
	f1()

	/*
		闭包是函数和环境变量的结合，i最后停留在2，所以这里打印出来的两个结果都是2
	*/
	for _, f2 := range closure2() {
		f2()
	}

	/*
		解决closure2问题的方法就是，要么环境变量不同，要么传参赋值
	*/
	// 环境变量不一样
	for _, f3 := range closure3() {
		f3()
	}
	// 赋值
	closure4()

}

/*
	延迟调用：
	defer向当前函数注册后执行的函数调用，这种调用被称为延迟调用。

	defer是FILO的
*/

func TestDefer(t *testing.T) {
	defer func() { fmt.Println("Defer 1 ...") }()
	defer func() { fmt.Println("Defer 2 ...") }()
	fmt.Println("Run")
}

func deferFunc() (z int) {
	defer func() {
		fmt.Println("Defer : ", z)
		z += 100
	}()

	z = 100
	return z
}

/*
	延迟调用不是Ret汇编，所以defer是可以改变返回值得
*/
func TestDeferNotRet(t *testing.T) {
	//Defer :  100
	//TestDeferNotRet :  200
	// 顺序是 z == 100 -> defer z += 100 -> ret汇编
	fmt.Println("TestDeferNotRet : ", deferFunc())
}

/*
	延迟调用需要注意不要放在循环里面，如果真想放到循环里面，需要封装函数
*/

func TestDeferLoop(t *testing.T) {
	do := func() {
		defer func() { fmt.Println("do finish") }() // 这里的defer是在do函数里面的，每次do结束都会执行
		fmt.Println("do begin")
	}

	for i := 0; i < 2; i++ {
		// 注意：这个defer是在TestDeferLoop函数里面的，他是在TestDeferLoop结束后一次运行多次
		defer func() { fmt.Println("TestDeferLoop do finish", i) }() //defer func也是一个闭包，所以i两次打印都是2
		do()
	}

	//do begin
	//do finish
	//do begin
	//do finish
	//TestDeferLoop do finish 2
	//TestDeferLoop do finish 2
}

/*
	延迟调用性能：
	延迟调用需要注册，调用等操作，还有额外的缓存开销
*/

var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}
func deferCall() {
	defer m.Unlock()
	m.Lock()
}

func BenchmarkCall(b *testing.B) {
	//BenchmarkCall-8   	81121734	        14.9 ns/op
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkDeferCall(b *testing.B) {
	//BenchmarkDeferCall-8   	63802018	        17.0 ns/op
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}

/*
	panic 和 recover 的使用：
	有点像try/catch，但是panic和recover并不是语句，他是方法
*/

func TestPanicAndRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // defer panic -> recover()只会捕获最后一个panic，有点类似rethrow
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("TestPanicAndRecover panic")
}
