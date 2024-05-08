package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello User, This is Calculator made by using Golang!")

	for {
		fmt.Println("\n1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")

		var choice int
		var first, second int

		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("\nExiting the Calculator")
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Println("\nInvalid Choice")
			continue
		}

		fmt.Print("\nEnter first number: ")
		fmt.Scanln(&first)

		fmt.Print("Enter second number: ")
		fmt.Scanln(&second)

		switch choice {
			case 1: 
				fmt.Println("\nAddition of",first ,"and" ,second ,"is: " ,add(first, second))
			case 2: 
				fmt.Println("\nSubtraction of",first ,"and" ,second ,"is: " ,sub(first, second))
			case 3: 
				fmt.Println("\nMultiplication of",first ,"and" ,second ,"is: " ,mul(first, second))
			case 4: 
				fmt.Println("\nDivision of",first ,"and" ,second ,"is: " ,div(first, second))
		}
	}
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}
