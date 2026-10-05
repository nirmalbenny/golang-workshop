package main

import (
	"fmt"
	"time"
)

func main() {
	tmp := 16
	if tmp > 18 {
		fmt.Println("Its Hot : ", tmp)
	} else {
		fmt.Println("Its cold")
	}
	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}
	fmt.Println(userAccess)

	if userAccess["jane"] {
		fmt.Println(userAccess["jane"])
	}
	if hasAccess, ok := userAccess["jane"]; hasAccess && ok {
		fmt.Println(hasAccess, ok)
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning...")
	case hour < 17:
		fmt.Println("Good afernoon")
	default:
		fmt.Println("Good Evening")
	}

	checkType := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("number")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("Unknwon type")
		}

	}
	checkType(123)
	checkType("wdwedwed")
	checkType(true)
	checkType(3434.3434)
}
