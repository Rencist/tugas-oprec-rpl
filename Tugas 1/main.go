package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
    length int
}

func insert(list *LinkedList, data int) {
    temp1 := &Node{data, nil}
    if list.head == nil {
        list.head = temp1
    } else {
        temp2 := list.head
        for temp2.next != nil {
            temp2 = temp2.next
        }
        temp2.next = temp1
    }
    list.length += 1
}

func (list *LinkedList) delete() {
    temp1 := list.head
    var temp2 *Node
    for temp1.next != nil {
        temp2 = temp1
        temp1 = temp1.next
    }
    temp2.next = nil
    list.length -= 1
}

func (list *LinkedList) show() {
	p := list.head
	for p != nil {
		fmt.Printf("%v ", p.data)
		p = p.next
	}
}

func main() {
    list := LinkedList{}
	insert(&list, 1)
	insert(&list, 2)
	insert(&list, 3)
	insert(&list, 4)
	insert(&list, 5)
	list.delete()
	list.show()
}