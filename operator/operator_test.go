package operator

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
	位运算
*/
func TestTypeOperator(t *testing.T) {
	var a, b int8 = 5, 3
	fmt.Printf("%.4b \n", a)    // 0101
	fmt.Printf("%.4b \n", b)    // 0011
	fmt.Printf("%.4b \n", a&b)  // 0001 与运算，都为1
	fmt.Printf("%.4b \n", a|b)  // 0111 或运算，至少一个1
	fmt.Printf("%.4b \n", a^b)  // 0110 亦或运算，只有一个1
	fmt.Printf("%.4b \n", ^a)   // 1010 按位取反
	fmt.Printf("%.4b \n", a&^b) // 0100 按位清楚
	fmt.Printf("%.4b \n", a<<2) // 10100 统一左移2位
	fmt.Printf("%.4b \n", b>>3) // 0000 统一右移3位
}

/*
	自增和自减，现在不是运算符了，只是用作独立句子，不能用于表达式
*/
func TestAutoIncreaseAndDelete(t *testing.T) {
	a := 1

	//++a, 不可以前置

	//if (a++)>1{}, 不可以用作表达式

	a++
	fmt.Println(a)
}

/*
	内存地址： 每个"字节单元"的唯一编号
	指针： 记录"内存地址"的实体
			| p:=&x 	| a := 100
	Memory	| 0x1200	| 100
	Address	| 0x800		| 0x1200

	Go 有两种指针
	1. 普通指针：这种指针是不可以运算的，但是可以做比较
	2. uintptr：他可以运算，但是在c里面，万能指针是可以安全持有对象的，但是在go里面，他只是一个普通的数据类型，他是可能被回收的。

	那么 普通指针 和 uintptr 是如何互换的呢？ 答： 通过unsafe.Pointer
	unsafe.Pointer 可以和 普通指针 进行相互转换；
	unsafe.Pointer 可以和 uintptr 进行相互转换。
*/
func TestPrtAndMemoryAddress(t *testing.T) {
	// & 和 * 的区别
	a := 100
	var p *int = &a    // 取地址运算符：&，获取实体的内存地址
	*p += 20           // 指针运算符： *, 通过指针间接引用实体
	fmt.Println(p, *p) //0xc000016338 120

	// Go是值传递，如果需要改的话，需要通过指针来实现引用传递
	person := &Person{
		Address: Address{Name: "GD"},
		Name:    "Leon",
		Age:     18,
	}
	UpdatePersonInfo(person, "SH")
	fmt.Println(person)

	// unsafe.Pointer的运用
	offsetof := unsafe.Offsetof(person.Age)          // Age 相对于 person的偏移量
	personUintptr := uintptr(unsafe.Pointer(person)) // person的unsafePointer转化为uintptr
	ageAddress := personUintptr + offsetof           // uintptr可以运算，相加就是Age的地址
	pointer := unsafe.Pointer(ageAddress)            // 指针地址
	*(*int)(pointer) = 20                            // 把unsafe.Pointer转化为*int指针，然后通过指针操作符*获取引用对象Age并且修改为20
	fmt.Println(person)                              //person的Age的确被修改了
}

/*
	初始化:
	对于符合类型（数组，切片，字典，结构体）变量初始化的时候的注意事项
*/

func TestInit(t *testing.T) {
	//数组
	array := []int{1, 2, 3}
	fmt.Println(array)

	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 1)
	slice = append(slice, 1)
	slice = append(slice, 1)
	slice = append(slice, 1)
	slice = append(slice, 1)
	fmt.Println(slice)

	m := map[string]int{"1": 1, "2": 2}
	fmt.Println(m)

	person := &Person{
		Address: Address{Name: "GD"},
		Name:    "Leon",
		Age:     18,
	}
	fmt.Println(person)
}

/*
	控制流： if
*/
func TestIf(t *testing.T) {
	// 基本用法：
	a := 1
	if a == 1 {
		fmt.Println(a)
	}

	// 支持数据初始化
	if b := a + 10; a > 10 {
		fmt.Println(b)
	} else {
		fmt.Println("a <= 10")
	}
}

/*
	控制流：switch
*/

func TestSwitch(t *testing.T) {
	// 基本用法
	a, b, c, d, e := 1, 2, 3, 2, 2

	switch d {
	case a, b: //多条件匹配
		fmt.Println(d)
	case c: // 单条件匹配
		fmt.Println(d)
	case 4: // 常量
		fmt.Println(d)
	default: // 默认
		fmt.Println("2")
	}

	// 相邻的空case不会造成多条件匹配
	switch d {
	case b: //隐式：case b: break
	case e:
		fmt.Println(b)
	}

	// fallthrough
	switch x := 5; x {
	default:
		println(x) // 编辑器保证不会提前执行default块
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough // 默认是break，如果想继续往下走，要用fallthrough，而且他不再匹配case的表达式
	case 6:
		x += 20
		fmt.Println(x)
		break // break之后，就不能再fallthrough了
		fallthrough
	case 7:
		x += 30
		fmt.Println(x)
	}
}

/*
	控制流： for
*/
func TestFor(t *testing.T) {
	// 方式1
	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}

	// 方式2
	nums := []int{1, 2, 3}
	for idx, num := range nums {
		fmt.Println(idx, num)
	}

	// 方式3
	i := 1
	for {
		if i == 2 {
			break
		}
		fmt.Println(i)
		i += 1
	}

	/*
		注意： for循环是会进行拷贝的，受直接影响的就是数组，可改用指针数组或者slice来解决
	*/
	data := [3]int{10, 20, 30}
	for i, x := range data {
		if i == 0 {
			data[0] = 100
			data[1] = 200
			data[2] = 300
		}
		/*
			x: 10, data: 100
			x: 20, data: 200
			x: 30, data: 300
			可以看出，data是被复制了一份，所以在循环里面你即便修改了data里面的值，x也不会被改变
		*/
		fmt.Printf("x: %d, data: %d \n", x, data[i])
	}
	fmt.Println(data) //[100 200 300]

	for i, x := range data[:] { // 这里只是复制了slice，没有复制底层的array
		if i == 0 {
			data[0] = 111
			data[1] = 222
			data[2] = 333
		}
		/*
			x: 100, data: 111 // 当i==0的时候，x已经取出来了
			x: 222, data: 222
			x: 333, data: 333
		*/
		fmt.Printf("x: %d, data: %d \n", x, data[i])
	}
}
