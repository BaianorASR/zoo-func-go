package main

import (
	"fmt"
	"zoologic/zoo"
)

func main() {
	results := zoo.GetSpeciesByIds("ef3778eb-2844-4c7c-b66c-f432073e1c6b", "0938aa23-f153-4937-9f88-4858b24d6bce")
	fmt.Println(results)
}
