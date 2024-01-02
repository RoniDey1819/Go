package main

import "fmt"

func main() {
	//act as while loop
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	//classical for loop
	for i := 100; i <= 105; i++ {
		fmt.Println(i)
	}
	//for without a condition will loop repeatedly until you
	// break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}
	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
