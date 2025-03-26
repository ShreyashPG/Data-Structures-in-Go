package main  

import (
	"fmt"
	"time"
)

func main(){
	
	// go sleepyGopher()
	// time.Sleep(4 * time.Second)

	// for i:= 0 ;i<5;i++{
	// 	go sleepyGopher(i)
	// }
	// time.Sleep(4 * time.Second) 


	//channels in go
	c:= make(chan int)
	for i:= 0 ;i<5;i++{
		go sleepyGopher(i, c)
	}
	for i:=0;i<5;i++{
		gopherID := <-c
		fmt.Println("gopher",gopherID,"has finished sleeping")
	}
	
}

// func sleepyGopher(id int){
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Gopher",id,"...snore...")

// }

func sleepyGopher(id int, c chan int){
	time.Sleep(3 * time.Second)
	fmt.Println("Gopher",id,"...snore...")
	c<-id

}