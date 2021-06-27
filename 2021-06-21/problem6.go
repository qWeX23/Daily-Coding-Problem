package main

import "fmt"

type LinkedList struct {
	head, tail *Node
}
type Node struct {
	prev, next *Node
	msg        string
}

func (ll *LinkedList) add(n *Node) {
	if ll.tail != nil {
		ll.tail.next = n
	}
	ll.tail = n

}

func main() {

	first := Node{
		msg:  "weeee",
		next: nil,
	}
	second := Node{
		prev: &first,
		msg:  "2",
		next: nil,
	}
	first.next = &second

	ll := LinkedList{
		head: &first,
		tail: &second,
	}
	third := Node{
		prev: nil,
		next: nil,
		msg:  "assdfewqet",
	}
	ll.add(&third)
	currentNode := *ll.head
	for {
		fmt.Println(currentNode)
		if currentNode.next == nil {
			break
		} else {
			currentNode = *currentNode.next

		}

	}
}
