package main

import "fmt"

func main() {
	const USDtoEUR = 0.89
	const USDtoRUB = 91.13
	const EURtoRUB = USDtoRUB / USDtoEUR
	fmt.Printf("Курс евро к рублю: %.2f", EURtoRUB)
	userInput := scanUserInput("Введите текст: ")
	fmt.Println("Вы ввели:", userInput)
}

func scanUserInput(prompt string) string {
	var input string
	fmt.Print("Введите число: ")
	fmt.Scan(&input)
	return input
}
