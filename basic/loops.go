package main

import "fmt"

func main() {
	fmt.Println("---LOOPS---")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// while loop style for
	k := 6
	for k > 0 {
		fmt.Println("K : ", k)
		k--
	}
	//---inifite loop
	counter := 0
	for {
		fmt.Println("Counter : ", counter)
		counter++
		if counter >= 5 {
			break
		}
	}
	fmt.Println("Odd Numbers Missing")

	for j := 0; j < 10; j++ {
		if j%2 != 0 {
			continue
		}
		fmt.Println("NoDe : ", j)
	}

	items := [3]string{"Hi", "HOw", "are you"}
	for i := 0; i < 3; i++ {
		fmt.Println(items[i])
	}

	for _, value := range items {
		fmt.Println("Value : ", value)
	}
}
