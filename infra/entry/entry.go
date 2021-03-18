package main

import (
	"fmt"
	"learn-go/infra"
)

/*
	2、 duck Typing 概念
	强实现：必须所有条件都符合，才算是鸭子
	弱实现：长得像鸭子的就是鸭子

	3、 接口的定义
	接口是定义在使用者身上的， java是定义在创作者身上。
	好处：因为很多时候你在写代码的时候想不了那么多，也做不到完美的接口。
		 而duck typing就很方便在后期维护和扩展
*/
type Reader interface {
	Download(url string) string
}

type Writer interface {
	Write(context string) bool
}

/*
	5、 组合接口
*/
type RWHandler interface {
	Reader
	Writer
}

/*
	1、 接口的概念
	抽象出不同对象和结构中相同的属性和行为
*/
func main() {

	reader := infra.HttpReader{}
	resp := reader.Download("http://www.baidu.com")
	fmt.Println(resp[:100])

	faker := infra.FakerReader{}
	fakerResp := faker.Download("http://www.baidu.com")
	fmt.Println(fakerResp)

	rw := infra.HttpRWHandler{}
	rw.Download("http://www.baidu.com")
	rw.Write("hi")
	fmt.Println(rw.Context)
}
