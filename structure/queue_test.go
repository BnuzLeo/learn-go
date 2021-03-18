package structure

import (
	"fmt"
	"testing"
)

type Queue []int

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Head() int {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *Queue) Pop() int {
	result := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return result
}

func TestQueue(t *testing.T) {
	q := Queue{1, 2, 3}
	q.Push(4)
	fmt.Println(q.Head())
	fmt.Println(q.Pop())
	fmt.Println(q)
}
