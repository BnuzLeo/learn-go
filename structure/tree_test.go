package structure

import (
	"fmt"
	"testing"
)

type Node struct {
	value       int
	left, right *Node
}

func createNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func (n *Node) prefix() {
	if n == nil {
		return
	}
	n.left.prefix() // 无论是指针还是值，go都是用.来获取值
	fmt.Println(n.value)
	n.right.prefix()
}

func (n *Node) setValue(value int) {
	n.value = value
}

func TestTree(t *testing.T) {
	node3 := createNode(3)
	node4 := createNode(4)
	node5 := createNode(5)
	node6 := createNode(6)
	node7 := createNode(7)

	node3.left = node4
	node4.right = node5
	node3.right = node6
	node6.left = node7
	node3.prefix()
	fmt.Println("--------更新节点7 to 8--------")
	node7.setValue(8)
	node3.prefix()
}
