//https://play.golang.org/p/9HXRNtqtTg4

package main

import (
	"fmt"
	"time"
	"reflect"
	//"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	os := "linux"
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	
	//a:= time.Now().Unix() // both need to be nano unix!
	a:= time.Now().UnixNano() 
	time.Sleep(2*time.Second)
	//
	//time.Sleep(1*time.Nanosecond)
	
	//unix nano could have gone disastrously wrong!
	
	//b:= int32(time.Now().UnixNano())
	
	//var b int32
	//b=int32(time.Now().UnixNano())
	
	
	
	
	
	var b int64
	b=time.Now().UnixNano()
	
	
	
	fmt.Println(" ",a,b)
	fmt.Println(reflect.TypeOf(a).String())
	fmt.Println(reflect.TypeOf(b).String())
	
	
	//by default is int
	//timesec:=2 // this gives ./prog.go:52:19: invalid operation: timesec * int64(time.Nanosecond) (mismatched types int and int64)
	var timesec int64 =2
	
	
	
	//if b-a == int64(2*time.Second){
	if b-a == timesec*1000000000*int64(time.Nanosecond){
	fmt.Println("lola inside",int64(2*time.Second))
	}
	fmt.Println("lola",b-a)
	fmt.Println("lola time",timesec*int64(time.Nanosecond))
	fmt.Println("lola",int64(2*time.Minute))
	
	
}
