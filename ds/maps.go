package main

import "fmt"

func main() {
	fmt.Println("Starting....\n")
	studentGrades := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	studentGrades["Alice"] = 99
	fmt.Printf("%+v\n", studentGrades)
	alice, ok := studentGrades["Alice"]
	fmt.Printf("%v. %v", alice, ok)

	key := "James"
	if value, ok := studentGrades[key]; ok {
		fmt.Printf(">>> %s : %+v <<<\n", key, value)
	}
	delete(studentGrades, key)
	fmt.Printf("%+v\n", studentGrades)
	test := map[string]int{}

}
