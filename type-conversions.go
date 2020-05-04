package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	highestInt16 := int16(32767)
	lowestInt16 := int16(-32768)
	fmt.Printf("%v\n", highestInt16+1) // 32767
	fmt.Printf("%v\n", lowestInt16-1)  // -32768
	fmt.Println(math.MaxInt16 == (math.MinInt16-1)) // false
	fmt.Println(math.MaxInt16 == lowestInt16-1) // true
	fmt.Println(string(rune(65)))
	fmt.Println("12345" + strconv.Itoa(12345))
	fmt.Println(fmt.Sprintf("%v", 12345))
}
