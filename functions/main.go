package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	sum1 := sumup(1, 2, 3)
	sum2 := sumup(1, numbers...)
	fmt.Println(sum1)
	fmt.Println(sum2)
}

// variatic function
func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
