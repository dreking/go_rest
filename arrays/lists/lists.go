package lists

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := []string{"A", "B", "C"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	slicedHobbies1 := hobbies[:2]
	fmt.Println(slicedHobbies1)
	// slicedHobbies1 := hobbies[0:2]

	slicedHobbies2 := slicedHobbies1[1:3]
	fmt.Println(slicedHobbies2)

	// merge list
	newHobbies := []string{"D", "E"}
	hobbies = append(hobbies, newHobbies...)

	courses := []string{"API", "REST"}
	courses[1] = "NEW"
	courses = append(courses, "DAMN")
	fmt.Println(courses)

	products := []Product{{id: "1", title: "p1", price: 12.0}, {id: "2", title: "p2", price: 10.0}}
	products = append(products, Product{id: "3", title: "p3", price: 15})
	fmt.Println(products)
}

func slicesIntro() {
	// Dynamic array of structs
	products := []Product{{id: "1", title: "p1", price: 12.0}, {id: "2", title: "p2", price: 10.0}}
	products = append(products, Product{})
	fmt.Println(products)

	// Simple array with lenght
	prices := [4]float64{12, 11, 10, 9}
	fmt.Println(prices)

	// slices
	// featuredPrices := prices[1:3] // start and end index excluding last index)
	// featuredPrices := prices[:3] // start at the begining to the index(excluding the last index)
	featuredPrices := prices[1:] // start at the specified index to the last index
	fmt.Println(featuredPrices)
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)

	// slice metadata
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	highlightedPrices = highlightedPrices[:3]
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// array assigning
	var productNames [4]string = [4]string{"Book"}
	productNames[2] = "Carpet"
	fmt.Println(productNames)
}
