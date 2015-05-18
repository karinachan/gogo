package main

import (
"fmt" 
"time"
)

func fib(call int) {	
	var m = 1
	var n = 1
	fmt.Println("Fib", call,": ", m)
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		var temp = n
		n = n + m
		m = temp
		fmt.Println("Fib",call,":", m)
	}
}

func main() {
	fib(1)
	fib(2)
}




