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

package main

import (
	"fmt"
	"time"
)


func Sentname(ch chan string,name string){
	ch <- name
}

func Reciver(ch chan string){
	resi := <-ch
	time.Sleep(2*time.Second)
	fmt.Println(resi)
}
func main(){
	fmt.Println("Go routine started")
	ch := make(chan string)
	go Sentname(ch, "Anjal")
	Reciver(ch)
	fmt.Println("Go routine finished")
}














