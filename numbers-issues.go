package main

import "fmt"

func divideForThree(number int) []int {

	var results []int

	for i := 1; i < 100; i++ {

		if number%i == 0 {

			results = append(results, i)

		}

	}

	return results

}

func PinOrPan(numberOne int, numberTwo int) {

	for i := 1; i < 100; i++ {

		if i%numberOne == 0 {

			fmt.Println("pan")

		}

		if i%numberTwo == 0 {

			fmt.Println("pin")

		} else {

			fmt.Println(i)

		}

	}

}

func main() {

	results := divideForThree(100)
	fmt.Println(results)

	PinOrPan(3, 5)

}
