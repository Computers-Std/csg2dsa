package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data: data,
	}
}

func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}
	if n.Data == nil {
		return "<empty>"
	}
	val, ok := n.Data.(string)
	if ok {
		return val
	}
	return "Unknown Type"
}

type LinkedList struct {
	First *Node
}

func NewLinkedList(n *Node) *LinkedList {
	return &LinkedList{
		First: n,
	}
}

func (list *LinkedList) Read(index int) interface{} {
	if list.First == nil {
		return nil
	}
	currentNode := list.First
	currentIndex := 0

	for currentIndex < index {
		currentNode = currentNode.Next
		if currentNode == nil {
			return nil
		}
		currentIndex++
	}
	return currentNode.Data
}

func (list *LinkedList) Search(value interface{}) int {
	if list.First == nil {
		return -1
	}
	currNode := list.First
	for currIndex := 0; currNode != nil; currIndex++ {
		// checks both Type and Value.
		if currNode.Data == value {
			return currIndex
		}
		currNode = currNode.Next
	}
	return -1
}

func (list *LinkedList) Insert(index int, value interface{}) {
	newNode := NewNode(value)

	// Insert at beginning
	if index == 0 {
		newNode.Next = list.First
		list.First = newNode
		return
	}

	// Insert at elsewhere
	currNode := list.First
	for i := 0; currNode != nil && i < index-1; i++ {
		currNode = currNode.Next
	}
	// If currNode is nil, the index was out of bounds
	if currNode != nil {
		newNode.Next = currNode.Next
		currNode.Next = newNode
	}
}

func (list *LinkedList) Delete(index int) {
	if list.First == nil || index < 0 {
		return
	}
	if index == 0 {
		list.First = list.First.Next
		return
	}
	currNode := list.First
	for i := 0; currNode != nil && i < index-1; i++ {
		currNode = currNode.Next
	}
	if currNode == nil || currNode.Next == nil {
		return
	}
	nodeToDelete := currNode.Next
	currNode.Next = nodeToDelete.Next
	nodeToDelete.Next = nil
}

// ex-1
func (list *LinkedList) String() string {
	if list.First == nil {
		return "<empty list>"
	}
	currNode := list.First
	var sb strings.Builder

	for currNode != nil {
		if val, ok := currNode.Data.(string); ok {
			sb.WriteString(val)
		} else {
			fmt.Fprintf(&sb, "%v", currNode.Data)
		}
		if currNode.Next != nil {
			sb.WriteString(" -> ")
		}
		currNode = currNode.Next
	}
	return sb.String()
}

// ex-3
func (list *LinkedList) LastNode() *Node {
	if list.First == nil {
		return nil
	}
	// Traverse until we find a node with no 'Next'
	currNode := list.First
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	return currNode
}

// ex-4
func (list *LinkedList) ReverseRecursive(node *Node) {
	if node == nil || node.Next == nil { // its a tail
		list.First = node // make it as the head
		return
	}
	list.ReverseRecursive(node.Next) // recurse to the end

	// Current State: A -> B -> C -> nil
	// We are currently at node B (popping off the stack)

	node.Next.Next = node
	// Intuition: "Take the node after me (C), and make its 'Next'
	// pointer look back at me (B)."
	//
	// Result: B -> C -> B (A temporary loop)

	node.Next = nil
	// Intuition: "Now that C is pointing at me, I can stop pointing
	// at C."
	//
	// Result: C -> B -> nil
}

func (list *LinkedList) Reverse() {
	var prev *Node
	curr := list.First
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	list.First = prev
}

// ex-5
func (list *LinkedList) DeleteInMiddle(node *Node) {
	if node == nil || node.Next == nil { // we cannot delete the tail
		return
	}
	next := node.Next
	node.Data = next.Data
	node.Next = next.Next
}

func main() {
	node1 := NewNode("once")
	node2 := NewNode("upon")
	node3 := NewNode("a")
	node4 := NewNode("time")
	node1.Next, node2.Next, node3.Next = node2, node3, node4

	list := NewLinkedList(node1)
	fmt.Println(list)

	list.Insert(4, "There lived a Ghost")
	fmt.Println(list, "(Inserted)")

	list.Delete(4)
	fmt.Println(list, "(Deleted)")

	fmt.Println(list.LastNode())

	list.ReverseRecursive(list.First)
	fmt.Println(list)

	list.Reverse()
	fmt.Println(list)

	list.DeleteInMiddle(node2)
	fmt.Println(list)
}
