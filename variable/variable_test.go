package variable

import (
	"fmt"
	"testing"
	"unsafe"
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

/*
	基本类型

	- 单位
	比特 ：计算机以二进制来存储和发送消息的，二进制的一位就是1比特，也叫做1bit
	字节： 字节和字符有关，中文一般2字节，英文一般1字节，字节也叫做Byte。 1 Byte = 8 bit
	千字节： 因为计算机只认识二进制，所以千字节是2的10次方，所以是1024 Byte。 1 KB = 1024 Byte
*/
func TestBaseType(t *testing.T) {
	// bool | 布尔值 | 1 byte | 默认：false
	var a bool
	fmt.Println(a)

	/*
		byte(uint8) | 字节型 | 1 byte | 默认值：0
		rune(int32) | 字符型 | 4 byte | 默认值：0
	*/
	var b byte
	fmt.Println(b)

	/*
		整数型：int/uint，默认根据目标平台决定是32位或者64位

		int 代表有符号的整数型：表示范围是：-2147483648到2147483648，即-2^31到2^31次方
		uint 代表没有符号的整数型: 表示范围是：2^32即0到4294967295

		int/uint     | 整数型 | 4,8 byte | 默认值：0
		int8/uint8   | 整数型 | 1 byte   | 默认值：0
		int16/uint16 | 整数型 | 2 byte   | 默认值：0
		int32/uint32 | 整数型 | 4 byte   | 默认值：0
		int64/uint64 | 整数型 | 8 byte   | 默认值：0
	*/
	var c1 int
	var c2 uint
	fmt.Println(c1, c2)
	var c3 int8
	var c4 uint8
	fmt.Println(c3, c4)

	/*
		float32   | 浮点型 | 4 byte  | 默认值 0.0
		float64   | 浮点型 | 8 byte  | 默认值 0.0
		complex64 | 复数型 | 8 byte  | 默认值 0.0
		complex128| 复数型 | 16 byte | 默认值 0.0
	*/
	var d1 float32
	var d2 float64
	fmt.Printf("%.1f,%.1f \n", d1, d2)
	var d3 complex64
	var d4 complex128
	fmt.Printf("%.1f,%.1f \n", d3, d4) //(0.0+0.0i),(0.0+0.0i)
}

/*
	类型转换：因为隐式转换带来的问题远大于他的好处，所以Go强制显示类型转换

	注意：指针，单向通道或者没有返回值得函数类型，需要用括号括起来
	如：
	（*int)(&x)
	(<-chan int)(x)
	(func())(x)

*/
func TestTypeChange(t *testing.T) {
	a := 255
	b := byte(a)

	var c interface{}
	c = 10
	d := c.(int)
	fmt.Println(b, d)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))

	fmt.Printf("%b \n", a)
	fmt.Printf("%b \n", b)
}
