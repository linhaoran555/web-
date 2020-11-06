package main

import (
	"fmt"
	"os"
)

func WriteFile(path string)  {
	f,err:=os.Create(path)
	if err !=nil {
		fmt.Println("err=",err)
		return
	}
	defer f.Close()
	buf:="don't communicate by sharing memory share memory by communicating"
	f.WriteString(buf)

}
func ReadFile(path string)  {
	f,err:=os.Open(path)
	if err!=nil {
		fmt.Println("err=",err)
		return
		}
	defer f.Close()
	buf:=make([]byte,1024*5)
	f.Read(buf)
	fmt.Printf("%s",buf)
}

func main()  {
	path :="./proveb.txt"
	WriteFile(path)
	ReadFile(path)


}