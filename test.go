package main

import (
	"fmt"
	"math"
)

func main() {

	x := 9
	y := 3

	r := math.Sqrt(float64(x)*float64(y) + float64(x)/float64(y))
	fmt.Println(int(r))

}
