package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	prices[1] = 9.99
	// prices[2] = 11.99
	// fmt.Println(prices) // error
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices) // [10.99 9.99 5.99]
	fmt.Println(prices)        // [10.99 9.99]

	prices = append(prices, 5.99)
	fmt.Println(prices) // [10.99 9.99 5.99]

	prices = prices[1:]
	fmt.Println(prices) // [9.99 5.99]

	prices = append(prices, 5.99, 12.99, 29.99, 100.99)
	fmt.Println(prices)

	dicountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, dicountPrices...) // ... takes out all the elements in that list
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)

// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)
// 	fmt.Println(prices[2])

// 	// featurePrices := prices[1:3]
// 	// featurePrices := prices[:3]
// 	// featurePrices := prices[1:]
// 	// featurePrices := prices[-1:] // error: minus index is not allowed in Go
// 	featurePrices := prices[1:4] // 4 is the highest bound
// 	fmt.Println(featurePrices)

// 	highlightPrices := featurePrices[:1]
// 	fmt.Println(highlightPrices)

// 	// slice don't copy the array, it's just a reference to a part of the array
// 	featurePrices[0] = 199.99
// 	fmt.Println(featurePrices)   // [199.99 45.99 20]
// 	fmt.Println(highlightPrices) // [199.99]

// 	fmt.Println(len(featurePrices), cap(featurePrices)) // 3 3
// 	// capacity includes the left elements behind(only to the right side)
// 	fmt.Println(len(highlightPrices), cap(highlightPrices)) // 1 3

// 	highlightPrices = highlightPrices[:3]
// 	fmt.Println(highlightPrices)                            // [199.99 45.99 20]
// 	fmt.Println(len(highlightPrices), cap(highlightPrices)) // 3 3
// }
