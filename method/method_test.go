package method

import (
	"fmt"
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
