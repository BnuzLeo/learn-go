package concurrent

var num int

func Counter() int {
	num++
	return num
}
