package main

/*
 Primer liste u go-u, radi obrnuto iz nekog razloga?
*/
import (
	"fmt"
)

// Node : osnova liste
type Node struct {
	br   int
	next *Node
}

// List : za pracenje
type List struct {
	head   *Node
	length int
}

func makeNode(br int) *Node {
	return &Node{br, nil}
}

func appendList(head, element *Node) *Node {
	if head == nil {
		return element
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = element
	return head
}

func printList(head *Node) {
	temp := head
	for temp != nil {
		defer fmt.Printf("%d ", temp.br)
		temp = temp.next
	}
	defer fmt.Println()
}

func printListReversed(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.br)
		temp = temp.next
	}
	fmt.Println()
}

func appendFront(head, element *Node) *Node {
	if head == nil {
		return element
	}
	element.next = head
	return element
}
func main2() {

	list := List{nil, 0}
	lista := &list
	var command, x int

	for ok := true; ok; ok = command > 0 {
		fmt.Println("Komanda: ")
		fmt.Scanln(&command)
		switch command {
		case 1:
			fmt.Scanln(&x)
			lista.head = appendList(lista.head, makeNode(x))
			break
		case 2:
			fmt.Scanln(&x)
			lista.head = appendFront(lista.head, makeNode(x))
			break
		case 3:
			printListReversed(lista.head)
			break
		case 4:
			printList(lista.head)
			break
		default:
			break
		}
	}

}
