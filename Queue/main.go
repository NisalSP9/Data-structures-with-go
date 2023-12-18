package main

import "fmt"

type queue struct {
	items []int
}

func (q *queue) enqueue(item int) {
	fmt.Println("Enqueue item", item)
	q.items = append(q.items, item)
}

func (q *queue) dequeue() {
	if len(q.items) > 0 {
		fmt.Println("Dequeue item", q.items[0])
		q.items = q.items[1:]
	} else {
		fmt.Println("Queue is empty !!!!")
	}
}

func main() {

	var myQueue queue

	myQueue.enqueue(1)
	fmt.Println(myQueue.items)
	myQueue.enqueue(3)
	fmt.Println(myQueue.items)
	myQueue.enqueue(5)
	fmt.Println(myQueue.items)

	myQueue.dequeue()
	fmt.Println(myQueue.items)
	myQueue.dequeue()
	fmt.Println(myQueue.items)
	myQueue.dequeue()
	fmt.Println(myQueue.items)
	myQueue.dequeue()
	fmt.Println(myQueue.items)

}
