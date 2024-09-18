package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)

	// 1. Pass by value
	// adultYears := getAdultYears(age) // copy value of variable age to function getAdultYears

	// 2. Pass by reference (avoid unnecessary value copies)
	// adultYears getAdultYears(agePointer)
	// fmt.Println(adultYears)

	// 3. Pass by reference (directly mutate values)
	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

// func getAdultYears(age *int) int{
// 	return *age - 18
// }

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
