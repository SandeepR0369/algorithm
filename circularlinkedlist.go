package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Data int
	Next *Node
}

func New(i int) *Node {
	n := Node{Data: i, Next: nil}
	return &n
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtBeginning(i int) error {
	if l == nil {
		return errors.New(" LinkedList Not initialzed")
	}
	data := New(i)
	if l.head != nil {
		temp := l.head
		l.head = data
		data.Next = temp
	} else {
		l.head = data
	}
	return nil
}

func (l *LinkedList) InsertAtEnd(i int) error {
	if l == nil {
		return errors.New(" LinkedList Not initialzed")
	}
	data := New(i)
	if l.head != nil {
		temp := l.head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = data
	} else {
		l.head = data
	}
	return nil
}

func (l *LinkedList) DeleteNode(i int) error {
	if l == nil {
		return errors.New(" LinkedList Not initialzed")
	}
	if l.head == nil {
		return errors.New(" LinkedList Empty")
	}
	if l.head.Data == i {
		fmt.Println("Deleting Data from Head ", l.head.Data)
		l.head = nil
		return nil
	}
	data := l.head
	for data.Next != nil {
		if data.Next.Data == i {
			temp := data.Next
			data.Next = temp.Next
			temp = nil
			return nil
		}
		data = data.Next
	}
	fmt.Println("Element not found in list ")
	return nil
}

func (l *LinkedList) Search(i int) error {
	if l == nil {
		return errors.New(" LinkedList Not initialzed")
	}
	if l.head == nil {
		return errors.New(" LinkedList Empty")
	}
	if l.head.Data == i {
		fmt.Println("Element Found ")
		return nil
	}
	data := l.head
	for data.Next != nil {
		if data.Next.Data == i {
			fmt.Println("Element Found ")
			return nil
		}
		data = data.Next
	}
	fmt.Println("Element not found in list ")
	return nil
}

func (l *LinkedList) display() error {
	if l == nil {
		return errors.New(" LinkedList Not initialzed")
	}

	if l.head != nil {
		data := l.head
		for data != nil {
			fmt.Println(data.Data)
			data = data.Next
		}
	} else {
		errors.New(" LinkedList Empty")
	}
	return nil
}

func (l *LinkedList) KthNodeFromEnd(i int) error {
	if l == nil {
		return errors.New(" LinkedList Not initialzed")
	}
	j := 1
	A := l.head
	B := l.head
	if l.head != nil {
		for j <= i {
			if B.Next == nil {
				return errors.New("LinkedList too small ")
			}
			B = B.Next
			j = j + 1
		}
		for B.Next != nil {
			A = A.Next
			B = B.Next
		}
		fmt.Println(" Kth Node From End of list is ", A.Data)

	} else {
		errors.New(" LinkedList Empty")
	}
	return nil
}

func main() {

	var l LinkedList

	data := make([]int, 0)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Number: ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
		}
		i, err := strconv.Atoi(text)
		if err != nil {
			break
		}
		data = append(data, i)

	}

	for _, i := range data {
		l.InsertAtBeginning(i)
	}
	l.display()
	e := l.KthNodeFromEnd(2)
	if e != nil {
		fmt.Println("Error occured ", e.Error())
	}

}
