package main

import (
	"fmt"
)

// /////e1
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Doubler[T Number](num T) T {
	return num * 2
}

// //////////e2
type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprintf("MyInt:%d", mi)
}

func Print[T Printable](r T) {
	fmt.Println(r)
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

type List[T comparable] struct {
	head *Node[T]
	end  *Node[T]
}

func (l *List[T]) Add(r T) {
	newNode := Node[T]{
		val:  r,
		next: nil,
	}
	if l.head == nil {
		l.head = &newNode
		l.end = l.head
	} else {
		l.end.next = &newNode
		l.end = &newNode
	}
}

func (l *List[T]) Insert(r T, p int) {
	newNode := &Node[T]{
		val:  r,
		next: nil,
	}
	if l.head == nil {
		l.head = newNode
		l.end = newNode
		return
	}
	if p <= 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}
	temp := l.head
	for i := 0; i < p-1; i++ {
		if temp.next == nil {
			temp.next = newNode
			l.end = temp.next
			return
		} else {
			temp = temp.next
		}
	}
	newNode.next = temp.next
	temp.next = newNode

	if l.end == temp {
		l.end = newNode
	}
}

func (l *List[T]) Index(r T) int {
	cur := l.head
	for i := 0; ; i++ {
		if cur == nil {
			return -1
		} else if cur.val == r {
			return i
		} else {
			cur = cur.next
		}
	}
}

func main() {
	fmt.Println(Doubler(10))
	fmt.Println(Doubler(11.2))
	var i MyInt = 20
	Print(i)

}
