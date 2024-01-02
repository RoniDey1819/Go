package main

import (
	"fmt"
	"math"
)

// We can declare variable here or inside main function
const s = "Roni"

func main() {
	//const s string = "Roni"
	fmt.Println(s)

	const n = 500000000 //A const statement can appear anywhere a var statement can.
	var m = 120000 / 20

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(m)

	fmt.Println(int64(d)) //Type conversion from LargeFloat to LargeInt

	const p = 500000000
	var q = 3e20 / p
	fmt.Println(q)

	fmt.Println(math.Sin(n))
}
