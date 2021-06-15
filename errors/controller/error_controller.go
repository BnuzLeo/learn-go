package main

import (
	"fmt"
	"learn-go/errors/service"
	"reflect"
	"strconv"
)

var (
	methodNum = "0"
)

func main() {
	for {
		printMethodList()
		fmt.Scanln(&methodNum)
		num, err := strconv.Atoi(methodNum)
		if err != nil {
			fmt.Println("输入有误，请重新输入")
			fmt.Println("------------------分割线------------------")
			continue
		}
		triggerMethod(num)
		fmt.Println("------------------分割线------------------")
	}
}

func triggerMethod(num int) {
	defer func() error {
		if err := recover(); err != nil {
			fmt.Println("方法不存在，请重新输入")
		}
		return nil
	}()
	errService := service.NewErrService()
	ptr := reflect.ValueOf(&errService).Elem()
	method := ptr.Method(num)
	results := method.Call([]reflect.Value{})
	fmt.Printf("%+v \n", results[0])
}

func printMethodList() {
	errService := service.NewErrService()
	ptr := reflect.TypeOf(&errService).Elem()
	for i := 0; i < ptr.NumMethod(); i++ {
		method := ptr.Method(i)
		fmt.Printf("input 【%v】 to exec method 【%v】\n", i, method.Name)
	}
	fmt.Println("please select the method:")
}
