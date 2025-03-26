package main

import (
	"bufio"

	"fmt"
	_ "strings"
    "os"
    
)

func main() {


fmt.Println("Hello This go file is for array and strings practice")

var s string
    s = "Hello This is Shreyash"
    fmt.Println(s);

    //cannot use same variable name again in go if declared using var [name] [type]
    // var s string

    fmt.Print("Enter your name : ")
    fmt.Scanln(&s)
    fmt.Println(s);

    //scan only reads first word after space it does not read next word , example input : Shreyash Ghanekar
    //output : Shreyash
    // var s string

    // fmt.Print("Enter your name : ")
    // fmt.Scanf("%s", &s);
    // fmt.Println(s);


    // fmt.Print("Enter your name : ")
    // fmt.Scanln( &s);
    // fmt.Println(s);

    // var s1 string
    
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter your surname")
    s1 , err :=reader.ReadString('\n')

    if err!=nil {
        fmt.Println(err);
    }
     fmt.Println(s1)

    fmt.Println("Enter your message ")
    message , _ , _:= reader.ReadLine()

        fmt.Println(string(message))

        for i := 0; i<10;i++ {
            fmt.Println(i)
        }
            
        for i := range(10) {
                    fmt.Println(i)
        }

        


            
}
