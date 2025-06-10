package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func JustCode() {
	// пример использования Reverse
	fmt.Println(Reverse("Hello World"))
}

func LearningScan() {
	var word1, word2 string
	fmt.Println("Введите два слова, разделенные пробелом:")
	fmt.Scan(&word1, &word2)
	fmt.Println("Вы ввели:", word1, word2)
}

func BufioExample() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите предложение:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Вы ввели:", input)
}

func Exam() {
	// Демонстрация арифметических операций
	var one int = 10
	var two float64 = 3.5

	fmt.Println("Арифметические операции:")
	fmt.Printf("Сумма: %.2f\n", float64(one)+two)
	fmt.Printf("Разность: %.2f\n", float64(one)-two)
	fmt.Printf("Произведение: %.2f\n", float64(one)*two)
	fmt.Printf("Частное: %.2f\n", float64(one)/two)

	// Расчет скорости
	fmt.Println("\nРасчет скорости")
	fmt.Println("Введите расстояние (км) и время (ч):")

	var distance float64
	var time int

	_, err := fmt.Scan(&distance, &time)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if time <= 0 {
		fmt.Println("Ошибка: время должно быть положительным числом")
		return
	}

	speed := distance / float64(time)
	fmt.Printf("Скорость: %.2f км/ч\n", speed)
}
