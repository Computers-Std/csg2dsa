package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Data string // interface{}
	Next *Node
	Prev *Node
}

func NewNode(data string) *Node {
	return &Node{
		Data: data,
	}
}

func (dlist *DoublyLinkedList) InsertAtEnd(value string) {
	newNode := NewNode(value)
	if dlist.First == nil {
		dlist.First = newNode
		dlist.Last = newNode
		return
	} else {
		newNode.Prev = dlist.Last // <-
		dlist.Last.Next = newNode // ->
		dlist.Last = newNode      // update last
	}
}

func (dlist *DoublyLinkedList) RemoveAtBegin() *Node {
	removedNode := dlist.First
	dlist.First = dlist.First.Next
	return removedNode
}

func (dlist *DoublyLinkedList) String() string {
	if dlist.First == nil {
		return "<empty list>"
	}
	currNode := dlist.First
	var sb strings.Builder
	for currNode != nil {
		sb.WriteString(currNode.Data)
		if currNode.Next != nil {
			sb.WriteString(" <-> ")
		}
		currNode = currNode.Next
	}
	return sb.String()
}

type DoublyLinkedList struct {
	First *Node
	Last  *Node
}

// ex-2
func (dlist *DoublyLinkedList) Reverse(node *Node) {
	if node == nil {
		// Swap the head and tail of the struct once finished
		dlist.First, dlist.Last = dlist.Last, dlist.First
		return
	}
	// Swap the individual node's pointers
	node.Next, node.Prev = node.Prev, node.Next
	// Since we swapped them, the "next" node is now in node.Prev
	dlist.Reverse(node.Prev)
}

func main() {
	node1 := NewNode("once")
	node2 := NewNode("upon")
	node3 := NewNode("a")
	node4 := NewNode("time")
	node1.Prev, node1.Next = nil, node2
	node2.Prev, node2.Next = node1, node3
	node3.Prev, node3.Next = node2, node4
	node4.Prev, node4.Next = node3, nil

	dlist := NewDoublyLinkedList(node1, node4)
	fmt.Println(dlist)

	dlist.Reverse(dlist.First)
	fmt.Println(dlist)
}

func NewDoublyLinkedList(first, last *Node) *DoublyLinkedList {
	return &DoublyLinkedList{
		First: first,
		Last:  last,
	}
}

type Queue struct {
	list *DoublyLinkedList
}

func NewQueue() *Queue {
	return &Queue{
		list: NewDoublyLinkedList(nil, nil),
	}
}

func (q *Queue) Enqueue(element string) {
	q.list.InsertAtEnd(element)
}

func (q *Queue) Dequeue() *Node {
	return q.list.RemoveAtBegin()
}

func (q *Queue) Read() string {
	if q.list.First == nil {
		return ""
	} else {
		return q.list.First.Data
	}
}
