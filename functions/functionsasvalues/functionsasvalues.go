package functionsasvalues

import "fmt"

type functionFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7, 8}
	fmt.Println(numbers)

	doubleNumbers := transformNumbers(&numbers, double)
	fmt.Println(doubleNumbers)

	tripleNumbers := transformNumbers(&numbers, triple)
	fmt.Println(tripleNumbers)

	transformerFn1 := getTransformFunction(&numbers)
	transformerFn2 := getTransformFunction(&moreNumbers)
	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

}

// Functions as values that calls other functions
func transformNumbers(numbers *[]int, transform functionFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Functions returning another function
func getTransformFunction(numbers *[]int) functionFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
