package infra

import "fmt"

type FakerReader struct {
}

/*
	4、 接口的值类型
	f是值传递，无论是*FakerReader其实也是指针地址的一个复制

	根据具体业务，如果FakerReader小的话，可以不用指针。
	如果FakerReader大的话，就需要用到指针了
*/
func (f FakerReader) Download(url string) string {
	return fmt.Sprintf("Hi %v , I am a faker reader", url)
}
