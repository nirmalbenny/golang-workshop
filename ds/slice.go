package main

import (
	"fmt"
	"slices"
)

func main() {

	names := make([]string, 5, 10)
	names[0] = "nirmal"
	names[4] = "hello"
	fmt.Printf("\n%+v | Len %d | Cap %d", names, len(names), cap(names))
	names = append(names, "A")
	names = append(names, "A")
	names = append(names, "A")
	names = append(names, "A")
	names = append(names, "A")
	fmt.Printf("\n%+v | Len %d | Cap %d", names, len(names), cap(names))

	names = append(names, "A")
	fmt.Printf("\n%+v | Len %d | Cap %d", names, len(names), cap(names))
	slices.Reverse(names)
	fmt.Printf("\n%+v | Len %d | Cap %d", names, len(names), cap(names))

}
