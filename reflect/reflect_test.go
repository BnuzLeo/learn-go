package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
	反射无非以下几种：
	Type , Value, Method 和 Construct
*/

/*
	Type:
*/

func TestReflectType(t *testing.T) {
	// 1. Typeof的基本用法
	var age Number = 100
	tx := reflect.TypeOf(age)
	fmt.Println(tx.Name(), tx.Kind()) // Number int

	// 2. Typeof中基类型和指针类型的区别
	tp := reflect.TypeOf(&age)
	fmt.Println(tp.Name(), tp.Kind())  // ptr
	fmt.Println("tx == tp:", tx == tp) // false

	// 3. Elem()可以返回指针，数组，切片，字典以及通道的基本类型
	/*
		打印结果：
		Address reflect.Address 0
		  Country string 0
		  Province string 16
		  Street string 32
		  Door string 48
		Name reflect.Content 64
		Age reflect.Number 80
		sex string 88
	*/
	var p Person = Person{
		Address: Address{
			Country:  "China",
			Province: "GuangDong",
			Street:   "Xiangzhou",
			door:     2102,
		},
		Name: "Leon",
		Age:  26,
		sex:  "Boy",
	}
	person := reflect.TypeOf(&p)
	if person.Kind() == reflect.Ptr {
		person = person.Elem() // 获取基本类型
	}
	for i := 0; i < person.NumField(); i++ {
		field := person.Field(i)
		fmt.Println(field.Name, field.Type, field.Offset)
		if field.Anonymous { // 判断是否是匿名字段
			for x := 0; x < field.Type.NumField(); x++ {
				structField := field.Type.Field(x)
				fmt.Println(" ", structField.Name, structField.Type, structField.Offset)
			}
		}
	}

	//4. 通过名称查找匿名字段
	address, _ := person.FieldByName("Country")
	fmt.Println(address.Name, address.Type) // Country string
	//5. 通过下标来查找匿名字段
	province := person.FieldByIndex([]int{0, 1})
	fmt.Println(province.Name, province.Type) // Province string
	//6. 通过Tag获取属性
	provinceJson := province.Tag.Get("json")
	fmt.Println(provinceJson) // Province
	//7. 获取方法
	var b B
	ptr := reflect.TypeOf(&b).Elem()
	for i := 0; i < ptr.NumMethod(); i++ {
		method := ptr.Method(i)
		fmt.Println(method.Name, method.Type)
	}
}

/*
	Value:
*/

func TestReflectValue(t *testing.T) {
	//1. Value的基本用法
	a := 100
	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()
	fmt.Println(va.CanAddr(), va.CanSet()) // false, false
	fmt.Println(vp.CanAddr(), vp.CanSet()) // true, true

	//2. 通过反射设置字段
	person := new(Person)
	personValue := reflect.ValueOf(person).Elem()
	provinceField := personValue.FieldByName("Province")
	fmt.Println(provinceField.CanSet(), provinceField.CanAddr()) // true, true
	if provinceField.CanSet() {
		provinceField.SetString("GD")
	}
	fmt.Println(person.Province) // GD

	//3. 编辑不可修改的字段
	doorField := personValue.FieldByName("door")
	fmt.Println(doorField.CanSet(), doorField.CanAddr()) // false, true
	if doorField.CanAddr() {
		*(*int)(unsafe.Pointer(doorField.UnsafeAddr())) = 2103
	}
	fmt.Println(person.door) // 2103

	//4. 类型推断和转化
	newPerson := Person{Name: "Leon"}
	personPtr := reflect.ValueOf(&newPerson)
	if personPtr.CanInterface() {
		fmt.Println(personPtr.CanInterface()) // true
		p, ok := personPtr.Interface().(*Person)
		if ok {
			fmt.Println(p.Name) //Leon
		}
	}

	//5. 判断接口是否为空
	var c interface{} = nil
	var d interface{} = (*int)(nil)
	fmt.Println(c == nil)                             // true
	fmt.Println(d == nil, reflect.ValueOf(d).IsNil()) // false , true
}

/*
	Method:
*/

func TestReflectMethod(t *testing.T) {
	//1. 反射的方法调用
	person := Person{Name: "Leon"}
	elem := reflect.ValueOf(&person).Elem()
	runMethod := elem.MethodByName("Run")
	in := []reflect.Value{
		reflect.ValueOf(30),
		reflect.ValueOf(100),
	}
	results := runMethod.Call(in)
	for _, result := range results {
		fmt.Println(result) //run fail
	}

	//2. 执行可变参数的方法
	duplicationIns := []reflect.Value{
		reflect.ValueOf(30),
		reflect.ValueOf(100),
		reflect.ValueOf(200),
	}
	runLongDurationMethod := elem.MethodByName("RunLongDuration")
	duplicationResults := runLongDurationMethod.Call(duplicationIns) // 这样是OK的
	for _, result := range duplicationResults {
		fmt.Println(result) //run fail
	}
	duplicationResults = runLongDurationMethod.CallSlice([]reflect.Value{ // 这样就更加方便，定制可变长度
		reflect.ValueOf(30),
		reflect.ValueOf([]int{100, 200}),
	})
	for _, result := range duplicationResults {
		fmt.Println(result) //run fail
	}

}

/*
	Construct:
*/

func TestReflectConstruct(t *testing.T) {
	// 1. 通过反射来创建对象
	person := Person{Name: "leon"}
	elem := reflect.TypeOf(person)
	p := reflect.New(elem)
	fmt.Println(p.Type())
}
