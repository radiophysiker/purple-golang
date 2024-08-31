package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Константы для конвертации валют
const USDtoEUR = 0.89
const USDtoRUB = 91.13
const EURtoRUB = USDtoRUB / USDtoEUR

// Справочник валют
var currencies = []string{"USD", "EUR", "RUB"}

func main() {
	//  Шаг 1: Вводим исходную валюту
	sourceCurrency := scanInputCurrency("Введите исходную валюту:", currencies)
	fmt.Println("Вы ввели:", sourceCurrency)

	// Шаг 2: Вводим сумму для конвертации
	amount := inputAmount("Введите сумму для конвертации")
	fmt.Println("Вы ввели:", amount)

	// Шаг 3: Вводим целевую валюту
	lastCurrencies := getLastCurrencies(sourceCurrency, currencies)
	targetCurrency := scanInputCurrency("Введите целевую валюту:", lastCurrencies)
	fmt.Println("Вы ввели:", targetCurrency)
	result := convert(sourceCurrency, targetCurrency, amount)
	fmt.Printf("Сумма %.2f", result)
}

// Функция для ввода и проверки валюты
func scanInputCurrency(prompt string, currencies []string) string {
	for {
		fmt.Println(prompt, strings.Join(currencies, ", "))
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка ввода", err)
			continue
		}
		// Приводим к верхнему регистру для удобства пользования приложения
		input = strings.ToUpper(strings.TrimSpace(input))
		if isValidCurrency(input, currencies) {
			return input
		} else {
			fmt.Println("Введена некорректная валюта")
		}
	}
}

// Функция для валидации валюты в соответствии со справочником
func isValidCurrency(currency string, currencies []string) bool {
	for _, c := range currencies {
		if currency == c {
			return true
		}
	}
	return false
}

// Функция для получения справочника оставшихся валют, исключая введенную целевую валюта
func getLastCurrencies(filteredCurrency string, currencies []string) []string {
	var filteredCurrencies []string
	for _, currency := range currencies {
		if currency != filteredCurrency {
			filteredCurrencies = append(filteredCurrencies, currency)
		}
	}
	return filteredCurrencies
}

// Функция для ввода и проверки суммы для конвертации
func inputAmount(prompt string) float64 {
	for {
		fmt.Println(prompt)
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка ввода", err)
		}

		amount, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Введено невалидное число")
			continue
		}
		if amount > 0 {
			return amount
		} else {
			fmt.Println("Введите положительное число")
		}
	}
}

func convert(sourceCurrency, targetCurrency string, amount float64) float64 {
	switch {
	case sourceCurrency == "USD" && targetCurrency == "EUR":
		return amount * USDtoEUR
	case sourceCurrency == "USD" && targetCurrency == "RUB":
		return amount * USDtoRUB
	case sourceCurrency == "EUR" && targetCurrency == "USD":
		return amount / USDtoEUR
	case sourceCurrency == "EUR" && targetCurrency == "RUB":
		return amount * EURtoRUB
	case sourceCurrency == "RUB" && targetCurrency == "USD":
		return amount / USDtoRUB
	case sourceCurrency == "RUB" && targetCurrency == "EUR":
		return amount / EURtoRUB
	default:
		return amount
	}
}
