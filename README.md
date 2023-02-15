# golang-linkedlist
Singly Linked-list menggunakan Golang

linked list terdiri node yang dihubungkan. node tersebut terdiri dari value yang disimpan pada node tersebut dan pointer menuju node selanjutnya. Ujung awal dari linked-list disimpan untuk membantu menandai suatu linked-list. 
```go
type node struct {
	data int
	next *node
}
type Linkedlist struct {
	size int
	head *node
}
```
untuk setiap pembuatan linked-list baru dapat dilakukan intialiasi nilai awal dari setiap varibel pada linked-list
```go
func LL_init(list *Linkedlist) {
	list.head = nil
	list.size = 0
}
```
berdasarkan latihan [Modul 1](https://github.com/godlixe/modul-go/blob/main/dasar-dasar-golang.md#fungsi-multiple-return)
1. menambahkan element di akhir linked list(push_back)
```go
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
```
2. fungsi untuk menghapus elemen di akhir linked-list(pop_back)
```go
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
```
3. fungsi untuk mencetak seluruh elemen dalam linked-list
```go
func (list *Linkedlist) print(){
	temp :=list.head
	for temp != nil{
		fmt.Printf("%d -> ",temp.data)
		temp = temp.next
	}
	fmt.Printf("\n")
}
```

<h2>I/O</h2>
<h3>fungsi main atau input kedalam linked-list</h3>

```go
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
```


<h3>Output dari operasi</h3>


```
5 ->
5 -> 99 ->
5 -> 99 -> 8 ->
5 -> 99 -> 8 -> 78 ->
5 -> 99 -> 8 ->
5 -> 99 ->
5 ->

```

<h1 align="center">Matur Nuwun</h1>
