package main

import "fmt"

func main() {
	value1, value2 := getIntegerValues()

	var option int

	fmt.Print("Select one action to do: \n")
	fmt.Print("1 -> Add numbers\n")
	fmt.Print("2 -> Subtract numbers\n")
	fmt.Print("3 -> Divide numbers\n")
	fmt.Print("4 ->  Multiply numbers\n")
	fmt.Scanf("%d ", &option)

	switch option {
	case 1:
		fmt.Print(value1 + value2)
	case 2:
		fmt.Print(value1 - value2)
	case 3:
		fmt.Print(value1 / value2)
	case 4:
		fmt.Print(value1 * value2)
	default:
		fmt.Printf("%d, is not valid", option)

	}
}

func getIntegerValues() (int, int) {

	var (
		a int
		b int
	)

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d ", &a)
	fmt.Print("Enter another number: ")
	fmt.Scanf("%d ", &b)
	return a, b
}
