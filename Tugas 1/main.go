package main

import "fmt"

type Node struct {
    data int
    next *Node
    head *Node
}

// type LinkedList struct {
//     head *Node
//     length int
// }

func insertFront(list *Node, data int) {
    temp1 := &Node{data, nil, list.head}
    if list.head == nil {
        list.head = temp1
    } else {
        temp1.next = list.head
        list.head = temp1
    }
}

func deleteFront(list *Node) {
    temp1 := list.head
    if list.head == nil {
        fmt.Println("Linked List doesn`t have any element")
    } else {
        list.head = temp1.next
    }
}

func (list *Node) insertBack(data int) {
    temp1 := &Node{data, nil, list.head}
    if list.head == nil {
        list.head = temp1
    } else {
        temp2 := list.head
        for temp2.next != nil {
            temp2 = temp2.next
        }
        temp2.next = temp1
    }
    // list.length += 1
}

func (list *Node) deleteBack() {
    temp1 := list.head
    var temp2 *Node
    if list.head == nil {
        fmt.Println("Linked List doesn`t have any element")
    } else {
        for temp1.next != nil {
            temp2 = temp1
            temp1 = temp1.next
        }
        temp2.next = nil
        // list.length -= 1
    }
}

func (list *Node) showAll() {
	p := list.head
    if list.head == nil {
        fmt.Println("Linked List doesn`t have any element")
    } else {
        for p != nil {
            fmt.Printf("%v ", p.data)
            p = p.next
        }
    }
}

func (list *Node) showFront() {
	p := list.head
    fmt.Printf("%v ", p.data)
}

func (list *Node) showBack() {
	p := list.head
    if list.head == nil {
        fmt.Println("Linked List doesn`t have any element")
    } else {
        for p.next != nil {
            p = p.next
        }
        fmt.Printf("%v ", p.data)
    }
}

func main() {
    list := Node{}
	list.insertBack(1)
	// insertBack(&list, 2)
	// insertBack(&list, 3)
	// insertBack(&list, 4)
	// insertBack(&list, 5) 
	// insertFront(&list, 0)
	// list.insert(1)
	// list.insert(2)
	// list.insert(3)
	// list.insert(4)
	// list.insert(5)
    // deleteFront(&list)
    // list.deleteBack()
	// list.showFront()
    // list.showBack()
    list.showAll()
}