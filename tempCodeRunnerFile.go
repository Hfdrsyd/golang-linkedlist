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