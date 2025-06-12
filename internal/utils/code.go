package utils

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
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

func Calculator() {

	var a, b float32

	fmt.Scan(&a, &b)

	fmt.Println("Сложение:", a+b)
	fmt.Println("Вычитание:", a-b)
	fmt.Println("Деление:", a/b)
	fmt.Println("Умножение:", a*b)
}

func Lines() {

	var word_1 string = "Hello"
	var word_2 string = "World"

	wordsTogether1 := word_1 + " " + word_2
	fmt.Println(wordsTogether1)

	wordsTogether2 := word_1 + word_2
	fmt.Println(wordsTogether2)

	wordsTogether3 := word_1 + " " + word_2
	fmt.Println(wordsTogether3)

	text := "Этот текст состоит из 74 байт и 43 символов"
	var lengthInBytes int = len(text)
	var lengthInRunes int = utf8.RuneCountInString(text)
	fmt.Println("len(text): ", lengthInBytes, "and utf8.RuneCountInString(text): ",
		lengthInRunes)

	// Если нужно вычислить длину строки, содержащей только английские буквы,
	// символы и цифры, достаточно использовать функцию len(),
	// так как каждый из этих символов занимает один байт.

	// Если требуется узнать длину строки, в которой присутствуют буквы других
	// языков, такие как русские или арабские, следует применять
	// функцию utf8.RuneCountInString() из пакета "unicode/utf8".

	var r rune = 'B' // 'B' имеет код 66 в Unicode
	fmt.Println("Input r + 1: ", r+1)
	fmt.Println("Input r - 1: ", r-1)

	var value_r rune = 'Y'
	fmt.Println((value_r))

	var r1 rune = 'A'    // 'A' -> 65 в Unicode
	var r2 rune = r1 + 1 // 'B' -> 66 в Unicode
	var r3 rune = r2 + 1 // 'C' -> 67 в Unicode

	fmt.Println(string(r1) + string(r2) + string(r3)) // выведет "ABC"

}

func GetValues() (int, string) {
	return 50, "golang"
}

func TestGetValues() {
	f, g := GetValues()
	fmt.Println("f =", f, "g =", g)

}
