package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	lenght int
}

func (l *linkedList) addNode(node *node) {
	currentHead := l.head
	node.next = currentHead
	l.head = node
	l.lenght++
}

func (l *linkedList) displayData() {
	fmt.Println()
	lenght := l.lenght
	currentHead := l.head
	for lenght > 0 {
		fmt.Printf("%d ", currentHead.data)
		currentHead = currentHead.next
		lenght--
	}
	fmt.Println()
}

func (l *linkedList) removeNodeByData(data int) {
	lenght := l.lenght
	var previouseHead *node
	currentHead := l.head
	for lenght > 0 {
		if currentHead.data == data && previouseHead == nil {
			fmt.Print("match found, remove data ", currentHead.data)
			l.head = l.head.next
			l.lenght -= 1
			return
		}
		if currentHead.data == data && previouseHead != nil {
			fmt.Print("match found, remove data ", currentHead.data)
			previouseHead.next = currentHead.next
			l.lenght -= 1
			return
		} else {
			previouseHead = currentHead
			currentHead = currentHead.next
			lenght--
		}
	}
	if lenght == 0 {
		fmt.Println("No match found for ", data)
	}
	fmt.Println()
}

func main() {
	var myLinkedList linkedList
	node1 := &node{data: 1}
	myLinkedList.addNode(node1)
	node2 := &node{data: 3}
	myLinkedList.addNode(node2)
	node3 := &node{data: 5}
	myLinkedList.addNode(node3)
	node4 := &node{data: 7}
	myLinkedList.addNode(node4)
	myLinkedList.displayData()

	myLinkedList.removeNodeByData(33)
	myLinkedList.displayData()

	myLinkedList.removeNodeByData(1)
	myLinkedList.displayData()

	myLinkedList.removeNodeByData(7)
	myLinkedList.displayData()

	myLinkedList.removeNodeByData(3)
	myLinkedList.displayData()

	myLinkedList.removeNodeByData(5)
	myLinkedList.displayData()

	myLinkedList.removeNodeByData(20)
	myLinkedList.displayData()

}
