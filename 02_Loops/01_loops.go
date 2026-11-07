package main

import "fmt"

func main() {

	// FOR LOOP

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// While LOOP
	whi := 1
	for whi < 1000 {
		whi += whi
	}
	fmt.Println(whi)

}
