package main

import "fmt"

type node struct {
	data int
	next *node
}
type Linkedlist struct {
	size int
	head *node
}

func LL_init(list *Linkedlist) {
	list.head = nil
	list.size = 0
}
func (list *Linkedlist) LL_push(value int) {//merupakan method dari suatu pointer linkedlist
	nwnode := &node{
		next: nil,
		data: value,
		}
	list.size++
	if list.head == nil {
		list.head = nwnode
	} else {
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = nwnode
	}
}
func (list *Linkedlist) print(){
	temp :=list.head
	for temp != nil{
		fmt.Printf("%d -> ",temp.data)
		temp = temp.next
	}
	fmt.Printf("\n")
}
// CurTemp :=list.head
// 	NextTemp :=list.head.next
// 	for NextTemp.next != nil{
// 		CurTemp=CurTemp.next
// 		NextTemp=NextTemp.next
// 	}
// 	CurTemp.next =nil
func (list *Linkedlist) pop()  {
	temp:=list.head
	if temp.next==nil {
		list.head=nil
		return
	}
	for temp.next.next != nil{
		temp=temp.next
	}
	temp.next =nil
}
func main() {
	var myList Linkedlist
	LL_init(&myList)
	myList.LL_push(5)
	myList.print()
	myList.LL_push(99)
	myList.print()
	myList.LL_push(8)
	myList.print()
	myList.LL_push(78)
	myList.print()
	myList.pop()
	myList.print()
	myList.pop()
	myList.print()
	myList.pop()
	myList.print()
	myList.pop()
	myList.print()
}