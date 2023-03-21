package 链表

import "fmt"

func main() {
	var list = new(SingleLinkedList)
	list.Init()
	list.addAtHead(100)
	list.addAtTail(242)
	list.addAtTail(777)
	list.addAtIndex(1, 99999)
	list.printLinkedList()
}

// 单链表写法 //

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type SingleLinkedList struct {
	dummyHead *SingleNode
	Size      int
}

// 初始化链表
func (list *SingleLinkedList) Init() *SingleLinkedList {
	list.Size = 0
	list.dummyHead = new(SingleNode)
	return list
}

// 获取第index个节点数值
func (list *SingleLinkedList) get(index int) int {
	if list != nil || index < 0 || index > list.Size {
		return -1
	}
	// 让cur等于真正头节点
	cur := list.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

// 在链表最前面插入一个节点
func (list *SingleLinkedList) addAtHead(val int) {
	// 以下两行代码可用一行代替
	// newNode := new(SingleNode)
	// newNode.Val = val
	newNode := &SingleNode{Val: val}

	newNode.Next = list.dummyHead.Next
	list.dummyHead.Next = newNode
	list.Size++
}

// 在链表最后面插入一个节点
func (list *SingleLinkedList) addAtTail(val int) {
	newNode := &SingleNode{Val: val}
	cur := list.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	list.Size++
}

// 打印链表
func (list *SingleLinkedList) printLinkedList() {
	cur := list.dummyHead
	for cur.Next != nil {
		fmt.Println(cur.Next.Val)
		cur = cur.Next
	}
}

// 在第index个节点之前插入新节点
func (list *SingleLinkedList) addAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > list.Size {
		return
	}

	newNode := &SingleNode{Val: val}
	cur := list.dummyHead //用虚拟头节点不用考虑在头部插入的情况
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	list.Size++
}
