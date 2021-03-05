package variable

import (
	"fmt"
	"testing"
)

/*
	变量定义：
	变量在计算机里面是一段或者多段用来存储数据的内存

	Go是静态语言，所以在声明变量的时候，总是有固定的变量类型、
	我们这能修改他的变量值，但是不能修改他的变量类型
*/

func TestVarDefinition(t *testing.T) {
	// 方法一：
	var a int // 初始化0
	fmt.Println(a)
	var y = false // 自动判断为bool类型
	fmt.Println(y)

	// 方法二：
	var (
		b, c int       // 同时声明
		e, f = 1, "hi" // 同时声明，且不同类型
	)
	fmt.Println(b, c, e, f)

	// 方法三：简短模式
	g, h := 1, "hi" // 用:=来声明，就不需要写var了
	fmt.Println(g, h)
}

/*
	常量：就是表示一些恒定不变的值
*/

func TestConst(t *testing.T) {
	const x, y int = 123, 0x22
	const s = "hello world"
	const c = 'I'
	const (
		i, f = 1, 0.123
		b    = false
	)
}

/*
	枚举：Go没有明确定义枚举enum，不过可以借助iota标识实现一套自增枚举

	扩展 - 数字常量不会分配储存空间，无须像变量那样通过地址来取值
*/

func TestEnum(t *testing.T) {
	const (
		x = iota
		y
		z
	)
	fmt.Println(x, y, z) // 0,1,2

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB, GB) // 1024 1048576 1073741824

	const (
		_, _ = iota, iota * 10 // 0, 0 * 10
		a, b                   // 1,1 * 10
		c, d                   // 2,2 * 10
	)
	fmt.Println(a, b, c, d) // 1 10 2 20

	const (
		e = iota
		f = 1 // 中断iota
		g = 100
		h = iota // 要重新使用必须手动声明，且数字累计
	)
	fmt.Println(e, f, g, h) // 0 1 100 3
}
