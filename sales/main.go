package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 40.2323,
	"MUG":    23.2,
	"BOOK":   23.22,
	"HAT":    37,
}

func calculateItemPrice(itemCode string) (float64, bool) {

	basePrice, found := productPrices[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf("\n- item %s (Sale! Original!: $%.2f, Sale Price : $%.2f)",
					originalItemCode,
					basePrice,
					salePrice)
				return salePrice, true
			}
		}
		fmt.Printf("\nItem : [%s] not found", itemCode)
		return 0, false
	}
	fmt.Printf("\n- item %s, Sale price : $%.2f", itemCode, basePrice)
	return basePrice, found
}

func main() {
	fmt.Println("-----SALES ORDER PROCESSOR -------")
	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}
	var subTotal float64
	fmt.Println("-------- Processing order -------")
	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subTotal += price
		}
	}
	fmt.Println("\n--------------\nTotal : ", subTotal)
}
