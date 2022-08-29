package main

import (
	"fmt"
)

/*
[a c] -> [b] -> [a b c d e]
*/
func main() {
	head := &Node{
		array: [5]byte{'a', 'b'},
		len:   2,
	}

	second := &Node{
		array: [5]byte{'b'},
		len:   1,
	}

	third := &Node{
		array: [5]byte{'a', 'b', 'c', 'd', 'e'},
		len:   5,
	}

	head.next = second
	second.next = third

	ul := NewUnrolledLinkedList(head, head.len+second.len+third.len)
	for i := 0; i < ul.len; i++ {
		fmt.Printf("%s ", string(ul.get(i)))
	}

	fmt.Println()
	ul.insert('n', 0)
	for i := 0; i < ul.len; i++ {
		fmt.Printf("%s ", string(ul.get(i)))
	}

	fmt.Println()
	ul.insert('x', 9)
	for i := 0; i < ul.len; i++ {
		fmt.Printf("%s ", string(ul.get(i)))
	}

	fmt.Println()
	ul.insert('x', 4)
	for i := 0; i < ul.len; i++ {
		fmt.Printf("%s ", string(ul.get(i)))
	}

	fmt.Println()
	ul2 := NewUnrolledLinkedList(nil, 0)
	ul2.insert('x', 0)
	for i := 0; i < ul2.len; i++ {
		fmt.Printf("%s ", string(ul.get(i)))
	}
}

/*
Question
Given a LinkedList, every node contains a array. Every element of the array is char
implement two functions
1. get(int index) find the char at the index
2. insert(char ch, int index) insert the char to the index
*/

type Node struct {
	next  *Node
	array [5]byte
	len   int
}

type UnrolledLinkedList struct {
	head *Node
	tail *Node
	len  int
}

func NewUnrolledLinkedList(head *Node, totalLen int) *UnrolledLinkedList {
	ul := &UnrolledLinkedList{}
	ul.head = head
	ul.len = totalLen
	// in case totalLen is not available
	count := 0
	curr := ul.head
	for curr != nil {
		count += curr.len
		curr = curr.next
	}
	ul.len = count
	// find tail
	tail := ul.head
	for tail != nil && tail.next != nil {
		tail = tail.next
	}
	ul.tail = tail
	return ul
}

func (ul *UnrolledLinkedList) get(index int) byte {
	if index < 0 || index >= ul.len || ul.len == 0 {
		return '-'
	}
	curr := ul.head
	for curr != nil {
		if index >= curr.len {
			index -= curr.len
			curr = curr.next
		} else {
			return curr.array[index]
		}
	}

	return '-'
}

//case 1: insert at the end -> create a new node, new node is tail
//case 2: insert normal -> move to next index
//case 3: insert normal but a full node -> move last index to a new node
func (ul *UnrolledLinkedList) insert(char byte, index int) {
	ul.len++
	curr := ul.head
	for curr != nil {
		if index >= curr.len {
			index -= curr.len
			curr = curr.next
			continue
		}
		if curr.len == 5 {
			newNode := &Node{len: 1, array: [5]byte{curr.array[4]}}
			curr.len--
			newNode.next = curr.next
			curr.next = newNode
			if newNode.next == nil {
				ul.tail = newNode
			}
		}
		curr.len++
		for i := curr.len - 1; i > index; i-- {
			curr.array[i] = curr.array[i-1]
		}
		curr.array[index] = char
		break
	}
	if curr == nil {
		newNode := &Node{len: 1, array: [5]byte{char}}
		if ul.head == nil {
			ul.head = newNode
			ul.tail = newNode
		} else {
			ul.tail.next = newNode
			ul.tail = newNode
		}
	}
}

func (ul *UnrolledLinkedList) delete(index int) {
	if index < 0 || index >= ul.len {
		return
	}

	prev := &Node{}
	prev.next = ul.head
	curr := ul.head

	for curr != nil {
		if index >= curr.len {
			index -= curr.len
		}else {
			if curr.len == 1 {
				prev.next = curr.next
				if curr == ul.tail {
					ul.tail = prev
				}
			}else {
				for i := index; i < curr.len-1; i++ {
					curr.array[i] = curr.array[i+1]
				}
				curr.len--
			}
			return
		}
		curr = curr.next
		prev = prev.next
	}

	ul.len--
}
