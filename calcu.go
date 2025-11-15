// package main
// import "fmt"

// func main(){
// 	var num1,num2,choice int

// 	for{

// 		fmt.Println("Enter your choice")
// 		fmt.Println("1.Add")
// 		fmt.Println("2.multiplication")
// 		fmt.Println("3.divition")
// 		fmt.Println("4.substraction")
// 		fmt.Println("5.Exit")
// 		fmt.Scanln(&choice)

// 		fmt.Println("Enter your first num :")
// 		fmt.Scanln(&num1)
// 		fmt.Print("Enter your Second num :")
// 		fmt.Scanln(&num2)
// 		switch choice{
// 		case 1:
// 			fmt.Printf("%d + %d = %d",num1,num2,num1+num2)
// 		case 2:
// 			fmt.Printf("%d * %d = %d",num1,num2,num1*num2)
// 		case 3:
// 			fmt.Printf("%d / %d = %d",num1,num2,num1/num2)
// 		case 4:
// 			fmt.Printf("%d - %d = %d",num1,num2,num1-num2)
// 		case 5:
// 			return
// 		default:
// 			fmt.Println("Invalid Operation")

// 		}

// 	}
// }

// package main
// import ("fmt"
// "errors")

// var num1,num2 int

// func div(a,b int) (int, error){
// 	if b==0{
// 		return 0,errors.New("canot divisible by zero")
// 	}

// 	return a-b,nil

// }

// func main(){

// 	fmt.Scanln(&num1)
// 	fmt.Scanln(&num2)

// 	d,err:=div(num1,num2)
// 	if err != nil{
// 		fmt.Println(err)
// 	}else{
// 		fmt.Println(d)
// 	}
// }

// package main

// import ("fmt"
// "time")

//  func main(){
// 	cha:=make(chan bool)
// 	go func (){
// 		fmt.Println("go routine started")
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("go routine finished")
// 		cha<-true
// 		}()
// 		<-cha
// }

// create a channel pass the name call that name in main function
// 7592890623

// package main

// import (
// 	"fmt"
// 	"time"
// )


// func Sentname(ch chan string,name string){
// 	ch <- name
// }

// func Reciver(ch chan string){
// 	resi := <-ch
// 	time.Sleep(2*time.Second)
// 	fmt.Println(resi)
// }

// func main(){
// 	fmt.Println("Go routine started")
// 	ch := make(chan string)
// 	go Sentname(ch, "Anjal")
// 	Reciver(ch)
// 	fmt.Println("Go routine finished")
// }


// ===================================Linear search in go DSA=========================

// package main 
// import "fmt"



// func Linearserch(arr []int,target int)(string){
// 	for x,i:=range arr{
// 		if i==target{
// 			return fmt.Sprintf("%d found at %d index in the array",target,x)
// 		}
// 	}
// 	return fmt.Sprintf("%d not found in the array",target)
// }

// func mian(){
	

// }

// ==========================================================
// package main 
// import "fmt"

// func main(){
// 	arr :=[10]int{1,3,4,6,7}
// 	arr1 :=[5]int{1,3,4,6,7}


// 	appe:=append(arr,arr1...)
// 	fmt.Println(appe)
// }





// =====================================linked list==========================================

// package main 
// import "fmt"


// type Node struct{
// 	data int 
// 	next *Node

// }

// type LinkedList struct{
// 	head *Node
// }


// func (l *LinkedList) addNode(val int){
// 	newNode:=&Node{data:val}
// 	if l.head==nil{
// 		l.head=newNode
// 	}else{
// 		current:=l.head
// 		for current.next!=nil{
// 			current=current.next
// 		}
// 		current.next=newNode
// 	}

// }

// func (l LinkedList) printList(){
// 	current:=l.head
// 	for current!=nil{
// 		fmt.Printf("%d->",current.data)
// 		current=current.next
// 	}

// }

// func main(){
// 	list:=LinkedList{}
// 	list.addNode(1)
// 	list.addNode(2)
// 	list.addNode(3)
// 	list.addNode(4)
// 	list.addNode(5)
// 	list.printList()

// }


// =========================================================


// arrya to an linked list

// package main 
// import "fmt"


// type Node struct{
// 	data int
// 	next *Node
// }

// func ConverArray(arr []int)*Node{
// 	if len(arr)==0{
// 		return nil
// 	}
// 	head:=&Node{data:arr[0]}
// 	current:=head
// 	for i:=1;i<len(arr);i++{
// 		current.next=&Node{data:arr[i]}
// 		current=current.next
// 	}
// 	return head
// }

// func Printlist(h *Node){
// 	for h!=nil{
// 		fmt.Printf("%d->",h.data)
// 		h=h.next
// 	}
// }
// func main(){
// 	arr:=[]int{1,2,3,4,5}
// 	head:=ConverArray(arr)
// 	Printlist(head)
// 	fmt.Print("nil")
// }


// length of an linkedlist and check an element in linked list

// package main
// import "fmt"

// type Node struct{

// 	data int 
// 	next *Node
// }
// type Head struct{
// 	head *Node
// }

// func (h *Head) add(val int){
// 	newNode := &Node{data: val}
// 	if h.head == nil{
// 		h.head = newNode
// 	}else{
// 		current := h.head
// 		for current.next != nil{
// 			current = current.next
// 		}
// 		current.next=newNode

// 	}
// }

// func (h *Head) length(){
// 	count := 0
// 	current := h.head
// 	for current != nil {
// 		count+=1
// 		current = current.next
// 	}
// 	fmt.Print(count)
// }


// func main(){
// 	list:=Head{}
// 	list.add(1)
// 	list.add(2)
// 	list.add(3)
// 	list.length()
// }
// ===============================doubly linked list================================



// package main
// import "fmt"

// type Node struct{
// 	data int
// 	prev *Node
// 	next *Node
// }

// type hnt struct{
// 	head *Node
// 	tail *Node
// }

// func (dll *hnt) add(val int) {
// 	new:=&Node{data:val}
// 	if dll.head==nil{
// 		dll.head=new
// 		dll.tail=new	

// 	}else{
// 		new.prev=dll.tail	
// 		dll.tail.next=new
// 		dll.tail=new
// 	}
// }

// func (dll *hnt) forward() {
// 	current := dll.head
// 	for current != nil {
// 		fmt.Printf("%d->", current.data)
// 		current = current.next
// 	}
// 	fmt.Print("nil\n")
// }
// func (dll *hnt) backward() {
// 	current := dll.tail
// 	for current != nil {
// 		fmt.Printf("%d->", current.data)
// 		current = current.prev
// 	}
// 	fmt.Print("nil\n")
// }

// func main(){
// 	list:=hnt{}
// 	list.add(1)
// 	list.add(2)
// 	list.add(3)
// 	list.forward()
// 	list.backward()

// }


// array to doubly linked list

// package main
// import "fmt"


// type Node struct{
// 	data int
// 	prev *Node
// 	next *Node
// }

// type hnt struct{
// 	head *Node
// 	tail *Node

// }

// func doubly(arr []int) *hnt{
// 	list:=&hnt{}
// 	for i:=0;i<len(arr);i++{
// 		new:=&Node{data:arr[i]}
// 		if list.head==nil{
// 			list.head=new
// 			list.tail=new
// 		}else{
// 			list.tail.next=new
// 			new.prev=list.tail
// 			list.tail=new
// 		}

// 	}
// 	return list
// }

// func (dll *hnt) forward (){
// 	current := dll.head
// 	for current !=nil {
// 		fmt.Printf("%d->",current.data)
// 		current=current.next
// 	}
// }

// func (dll *hnt) backward(){
// 	current := dll.tail

// 	for current !=nil {
// 		fmt.Printf("%d->",current.data)
// 		current=current.prev
// 	}
// }

// func main(){
// 	arr:=[]int{1,2,3,4,5}
// 	list:=doubly(arr)
// 	list.forward()
// 	fmt.Print("nil")
// 	list.backward()
// 	fmt.Print("nil")

// }


// deleting an head node form singly linked list


// package main 
// import "fmt"

// type Node struct{
// 	data int
// 	next *Node
// }

// type head struct{
// 	head *Node
// }

// func (h *head) add(val int){
// 	new:=&Node {data:val}
// 	if h.head==nil{
// 		h.head=new
// 	}else{
// 		current:=h.head
// 		for current.next!=nil{
// 			current=current.next
// 		}
// 		current.next=new
// 	}
// }
// func (h *head) delete(){
// 	if h.head!=nil{
// 		h.head=h.head.next
// 	}
// }
// func (h *head) print(){
// 	current:=h.head
// 	for current!=nil{
// 		fmt.Printf("%d->",current.data)
// 		current=current.next
// 	}
// 	fmt.Println("ending")
// }

// func main(){
// 	list:=head{}
// 	list.add(1)
// 	list.add(2)
// 	list.add(3)
// 	list.add(4)
// 	list.add(5)
// 	list.print()
// 	fmt.Println("after deleting the head node")
// 	list.delete()
// 	list.print()
// }


// -------------------deleting the head from doubly linked list=====================

package main
import "fmt"

type Node struct{
	data int
	prev *Node 
	next *Node
}

type head struct{
	head *Node
	tail *Node
}

func (h *head) add(val int){
	new:=&Node{data:val}
	if h.head==nil{
		h.head=new
		h.tail=new
	}
	current:=h.head
	for current.next!=nil{
		current=current.next
	}
	current.next=new
	new.prev=current
	h.tail=new
}

func (h *head) delete(){
	if h.head!=nil{
		h.head=h.head.next
		h.head.prev=nil
	}
}

func (h *head) print(){
	current:=h.head
	for current!=nil{
		fmt.Printf("%d->",current.data)
		current=current.next
	}
	fmt.Println("ending")
}

func main(){
	list:=head{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)
	list.print()
	fmt.Println("after deleting the head node")
	list.delete()
	list.print()
}