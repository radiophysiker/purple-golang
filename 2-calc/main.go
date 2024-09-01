package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var operations = []string{"AVG", "SUM", "MED"}

func main() {
	//  Шаг 1: Вводим операцию
	operation := scanInputOperation("Введите операцию:")
	//  Шаг 2: Вводим числа
	numbers := scanInputNumbers("через запятую (2, 10, 9)")
	result, err := calculate(numbers, operation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Полученный результат: %.2f", result)
}

// Функция для ввода и проверки допустимых операций
func scanInputOperation(prompt string) string {
	for {
		fmt.Println(prompt, strings.Join(operations, ", "))
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка ввода", err)
			continue
		}
		// Приводим к верхнему регистру для удобства пользования приложения
		input = strings.ToUpper(strings.TrimSpace(input))
		if isValidOperations(input) {
			return input
		} else {
			fmt.Println("Введена некорректная операция")
		}
	}
}

func isValidOperations(operation string) bool {
	for _, o := range operations {
		if operation == o {
			return true
		}
	}
	return false
}

func scanInputNumbers(prompt string) []float64 {
	for {
		fmt.Println(prompt)
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка ввода чисел", err)
			continue
		}
		numbers, err := parseNumbers(input)
		if err != nil {
			continue
		}
		return numbers
	}
}

func parseNumbers(input string) ([]float64, error) {
	// Очищаем от лишних пробелов
	input = strings.TrimSpace(input)
	strNumbers := strings.Split(input, ",")
	numbers := make([]float64, 0, len(strNumbers))
	for index, strNum := range strNumbers {
		if strNum == "" {
			continue
		}
		strNum = strings.TrimSpace(strNum)
		num, err := strconv.ParseFloat(strNum, 64)
		if err != nil {
			fmt.Println(index+1, "число введено некорректно")
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func calculate(numbers []float64, operation string) (float64, error) {
	switch operation {
	case "AVG":
		return average(numbers), nil
	case "SUM":
		return sum(numbers), nil
	case "MED":
		return average(numbers), nil
	default:
		return 0, errors.New("неподдерживаемая операция")
	}
}

// Функция для вычисления среднего значения для массива чисел
func average(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

// Функция для вычисления суммы для массива чисел
func sum(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Функция для вычисления среднего значения для массива чисел
func median(numbers []float64) float64 {
	// сортируем массив для вычисления медианы
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 0 {
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}
