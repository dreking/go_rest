package maps

import "fmt"

func main() {
	// defining maps
	websites := map[string]string{"google": "https://google.com", "aws": "https://aws.com"}
	fmt.Println(websites)

	// getting value from key
	fmt.Println(websites["aws"])

	// assigning new key-value
	websites["linkedin"] = "http://linkedin.com"
	fmt.Println(websites)

	// update a key
	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	// delete a key
	delete(websites, "linkedin")
	fmt.Println(websites)
}
