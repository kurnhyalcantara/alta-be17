package main

import "fmt"

func main()  {
	prime := [5]int{2, 3, 5, 7, 11}

	// technique 1
	for index := 0; index < len(prime); index++ {
		fmt.Println(prime[index])
	}

	// technique 2
	for index, element := range prime {
		fmt.Println(index, "=>" , element)
	}
}