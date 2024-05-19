package main

import (
	"fmt"
	"math"
)

type SingleNode struct {
	value any
	next  *SingleNode
}

type DLLNode struct {
	value    any
	next     *DLLNode
	previous *DLLNode
}
type DoubleLinkedList struct {
	head *DLLNode
	tail *DLLNode
}

//var (
//	head SingleNode = SingleNode{value: 0, next: &n1}
//	n1   SingleNode = SingleNode{value: 1, next: &n2}
//	n2   SingleNode = SingleNode{value: 2, next: nil}
//
//)

func (dll *DoubleLinkedList) insertNode(value any, previous *DLLNode) {
	var newNode DLLNode = DLLNode{value: value, previous: previous, next: previous.next}
	if previous == dll.tail {
		dll.tail = &newNode
	}
	previous.next = &newNode
}

func (dll *DoubleLinkedList) createNode(value any) {
	var newNode DLLNode = DLLNode{value: value, previous: nil, next: dll.head}
	dll.head.previous = &newNode
	dll.head = &newNode
}

func (dll *DoubleLinkedList) printDLL() {
	current := dll.head
	for {
		if current.next != nil {
			fmt.Print(current, "->")
			current = current.next
			continue
		}
		fmt.Println(current)
		break
	}
}

func (dll *DoubleLinkedList) deleteNode(node *DLLNode) {
	// var previous DLLNode = node.previous
	if node == dll.head {
		fmt.Println("Deleting Head: ", dll.head)
		dll.head = node.next

		node.next.previous = node.previous
		return
	}
	if node == dll.tail {
		fmt.Println("Deleting Tail:", dll.tail)
		dll.tail = node.previous
		node.previous.next = node.next
		return
	}
	fmt.Print("Deleting Node:", node)
	node.previous.next = node.next
	node.next.previous = node.previous
}

func (s *SingleNode) getNext() *SingleNode {
	if s.next != nil {
		return s.next
	}
	fmt.Println("Node doesn't have next")
	return &SingleNode{}
}

func (s *SingleNode) deleteNext() {
	var nextNext SingleNode = *s.getNext().getNext()
	s.getNext().next = nil
	s.next = &nextNext
}

func (head *SingleNode) printSLL() {
	current := head
	for {
		if current.next != nil {
			fmt.Print(current, "-> ")
			current = current.next
			continue
		}
		fmt.Println(current)
		break
	}
}

func linearSearch(breaks []bool, ind, jump int) int {
	hi := len(breaks)
	fmt.Println("linearSearch, ", ind)
	for i2 := ind; i2 < hi; i2++ {
		fmt.Println("Ind2: ", i2, ind)
		if breaks[i2] == true {
			return ind + (i2 - ind)
		}
	}
	return -1
}

func CrystalBalls(breaks []bool) int {
	var hi int = len(breaks)
	var jump int = int(math.Sqrt(float64(hi)))
	ind := 0

	for {
		fmt.Println("Ind: ", ind, hi)
		if ind > hi {
			break
		}
		var value bool = breaks[ind]
		if value == false {
			if ind+jump < hi {
				ind += jump
				continue
			}
			return linearSearch(breaks, hi-jump, jump)
		}
		return linearSearch(breaks, ind-jump, jump)

	}

	return -1
}

func main() {
	//ballBreak := make([]bool, 10)
	//for i := 7; i < len(ballBreak); i++ {
	//	ballBreak[i] = true
	//}

	// fmt.Println(ballBreak)
	// index := CrystalBalls(ballBreak)
	// fmt.Println("CrystalBalls: ", index)

	var dll DoubleLinkedList
	var dllHead DLLNode = DLLNode{value: "head", next: nil, previous: nil}
	dll.head = &dllHead
	dll.tail = &dllHead
	dll.printDLL()
	dll.createNode("new head")
	dll.printDLL()
	dll.insertNode("inserted node", dll.head.next)
	dll.printDLL()
	dll.deleteNode(dll.head)
	dll.printDLL()
}
