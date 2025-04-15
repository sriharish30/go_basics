package maps

import "fmt"

func main() {
	//Declaring maps

	websites := map[string]string{
		"Google":             "http://Google.com",
		"Amazon Web Service": "http://Aws",
	}
	fmt.Println(websites)

	//print the values by calling the key

	fmt.Println(websites["Google"])

	// adding values in the map

	websites["LinkedIn"] = "http://linkedin.com"
	fmt.Println(websites)

	//deleting values in the map

	delete(websites, "Google")
	fmt.Println(websites)
}
