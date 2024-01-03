package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// courseRatings := map[string]float64{}
	courseRatings := make(floatMap, 2)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.7
	courseRatings.output()

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value", value)
	}
}
