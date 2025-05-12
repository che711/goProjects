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
	// простой однострочный комментарий

	/*
		Для чего нужны комментарии:
			Для объяснения сложных участков кода.
			Для описания функций и их назначения.
			Для временного отключения кода при отладке.
	*/

	//fmt.Print("Many\n")
	//fmt.Print("Thanks\n")
	//fmt.Print("To")
	//fmt.Print("All")
	//fmt.Print("My")
	//fmt.Print("Students")
	//
	//fmt.Print("\n")
	//fmt.Println("Hello", "my", "dear", "friends")

	/*
		// 1 вариант
		const a = "строка"
		//a = 'test_msg'
		fmt.Println(a)

		// 2 вариант
		var b, c, d string // пустая строка ""
		fmt.Println(b, c, d)

		// 3 вариант
		f, e := "new_test", "test"
		fmt.Println(e)
		fmt.Println(f)
	*/

	var text string = "string"
	fmt.Println(text)

	var b int = 123123123
	fmt.Println(b)
	var c int8 = 123
	fmt.Println(c)
	var d int16 = 1234
	fmt.Println(d)
	var e int32 = 123123123
	fmt.Println(e)
	var f int64 = 123123123123
	fmt.Println(f)

	var g uint = 999
	fmt.Println(g)
	var h uint8 = 100
	fmt.Println(h)
	var i uint16 = 1000
	fmt.Println(i)
	var j uint32 = 10000
	fmt.Println(j)
	var k uint64 = 10000000
	fmt.Println(k)

	var l float32 = 1000.0
	fmt.Println(l)
	var m float64 = 9999.123123123
	fmt.Println(m)

	var n complex64 = 100 + 100i
	fmt.Println(n)
	var o complex128 = 1000 + 1000i
	fmt.Println(o)

	var yes bool = true
	fmt.Println(yes)
	var no bool = false
	fmt.Println(no)
	fmt.Println("\n")

	someEmojis := "😀 😃 😄 😁 😆 😅 😂 🤣"
	someHieroglyphs := "𓀀 𓀁 𓀲 𓀕 𓀖 𓀗 𓀘 𓀙 "
	hieroglyph := '𓀙' // одинарные кавычки!!

	fmt.Println(someEmojis)
	fmt.Println("\n")
	fmt.Println(someHieroglyphs)
	fmt.Println("\n")
	fmt.Println(hieroglyph)
	fmt.Println("\n")

	// Базовые (встроенные) типы данных
	basicTypes := []interface{}{
		// Логический тип
		false,
		// Целочисленные (со знаком)
		int(0), int8(0), int16(0), int32(0), int64(0),
		// Целочисленные (без знака)
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0),
		// Числа с плавающей точкой
		float32(0.0), float64(0.0),
		// Комплексные числа
		complex64(0 + 0i), complex128(0 + 0i),
		// Строки
		"",
		// Символ (руна)
		rune('a'),
		// Псевдоним byte (то же, что uint8)
		byte('a'),
	}

	// Выводим типы
	fmt.Println("=== Базовые типы данных в Go ===")
	for _, v := range basicTypes {
		fmt.Printf("%T\n", v)

	}

	// Составные и другие типы (их нельзя просто создать в списке)
	fmt.Println("\n=== Составные и другие типы ===")
	fmt.Println("- Массивы (например, [3]int)")
	fmt.Println("- Срезы ([]int)")
	fmt.Println("- Структуры (struct { ... })")
	fmt.Println("- Указатели (*int)")
	fmt.Println("- Функции (func())")
	fmt.Println("- Интерфейсы (interface{})")
	fmt.Println("- Каналы (chan int)")
	fmt.Println("- Отображения (map[string]int)")

	fmt.Println("\n")
	var symbol rune = 'A'
	fmt.Println(symbol)
}

func LearningScan() {
	var word1, word2 string

	fmt.Println("Введите два слова, разделенные пробелом:")

	/*
		cчитывает два слова разделенные пробелом и сохраняет их в переменные
		cимвол амперсанд (&) перед переменной возвращает ее текущий адрес в оперативной памяти ->
		чтобы сохранить пользовательский ввод в переменную X, нужно передать ее в функцию fmt.Scan() с амперсандом (&) перед ней.
	*/

	fmt.Scan(&word1, &word2)

	fmt.Println("Вы ввели:", word1, word2)
}

func Bufio() {
	// Создаем новый reader чтобы считывать строки из консоли
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите предложение:")

	// Считываем строку до символа первого нажатия Enter
	input, _ := reader.ReadString('\n')

	// Выводим считанную строку
	fmt.Println("Вы ввели:", input)
}

func Exam() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
