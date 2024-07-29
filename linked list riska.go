package main

import "fmt"

// Node struct mewakili elemen dalam linked list
type Node struct {
    data int
    next *Node
}

// LinkedList struct untuk mengelola operasi pada linked list
type LinkedList struct {
    head *Node
}

// Insert menambahkan elemen baru di akhir linked list
func (ll *LinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

// Display mencetak semua elemen dalam linked list
func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    ll := &LinkedList{}
    ll.Insert(10)
    ll.Insert(20)
    ll.Insert(30)
    ll.Display()  // Output: 10 -> 20 -> 30 -> nil
}