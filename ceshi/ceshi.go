package main

import "fmt"



func Receiver(v interface{ })  {

	fmt.Printf("这个是%T",v)
}


func main()  {


	Receiver(34)


}