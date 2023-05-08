package main

import "fmt"

func main()  {
	var countries [2]string
	countries[0] = "Indonesia"
	countries[1] = "Malaysia"

	fmt.Println(countries)

	prime := [5]int{2,3,5,7,11}
	fmt.Println(prime)
	fmt.Println(len(prime))

	odd_numbers := [...]int{2,4,6,8,10,12,14}
	fmt.Println(odd_numbers)
	fmt.Println(len(odd_numbers))
}

