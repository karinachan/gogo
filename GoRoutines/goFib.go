package main

import (
"fmt" 
"time"
)

func f(b bool) {
	if b == true {
		fibIterative()	
	} else {
		fibRecursive(1,1,0)	
	}
}

func fibIterative() {	
	var m = 1
	var n = 1
	fmt.Println("FibIterative:", m)
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		var temp = n
		n = n + m
		m = temp
		fmt.Println("FibIterative:", m)
	}
}

func fibRecursive(m int, n int, stop int) {
	time.Sleep(100 * time.Millisecond)
	if stop <= 20 {
		fmt.Println("FibRecursive:", m)
		fibRecursive(n, m+n, stop + 1)
	}
}

func main() {
	go f(true)
	f(false)
}