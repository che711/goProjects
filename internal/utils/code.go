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
	fmt.Println('\n')

	someEmojis := "😀 😃 😄 😁 😆 😅 😂 🤣"
	someHieroglyphs := "𓀀 𓀁 𓀲 𓀕 𓀖 𓀗 𓀘 𓀙 "
	hieroglyph := '𓀙' // одинарные кавычки!!

	fmt.Println(someEmojis)
	fmt.Println('\n')
	fmt.Println(someHieroglyphs)
	fmt.Println('\n')
	fmt.Println(hieroglyph)
	fmt.Println('\n')

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

	fmt.Println('\n')
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
	//var name string
	//fmt.Scan(&name)
	//fmt.Println("Привет, ", name)

	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	//fmt.Println(input)

	//one, two, three, four = bufio.NewReader(os.Stdin)
	//fmt.Print(four, three, two, one)

	//var a, b, c, d int
	//fmt.Scan(&a, &b, &c, &d)
	//fmt.Println(d, c, b, a)

	//var word_a string
	//var word_b string
	//var word_c string
	//var word_d string
	//var word_e string
	//fmt.Scan(&word_a, &word_b, &word_c, &word_d, &word_e)
	//fmt.Println(word_e)
	//fmt.Println(word_d)
	//fmt.Println(word_c)
	//fmt.Println(word_b)
	//fmt.Println(word_a)

	var one int = 10
	var two float64 = 3.5

	// Преобразуем целое число `а` в число с плавающей точкой
	sum := float64(one) + two
	difference := float64(one) - two
	product := float64(one) * two
	quotient := float64(one) / two

	fmt.Println(5 % 2)
	fmt.Println("Cумма:", sum)
	fmt.Println("Разность:", difference)
	fmt.Println("Произведение:", product)
	fmt.Println("Частное:", quotient)

	//a := 10
	//
	//// порядок операции важен. ++a, --a и a- не существуют в Go.
	//a--
	//fmt.Println(a)
	//fmt.Println(-a)
	//
	//a++
	//fmt.Println(a)
	//fmt.Println(-a)
	//
	//	var a, b float32
	//	fmt.Scan(&a, &b)
	//
	//	fmt.Println("Сумма: ", a+b)
	//	fmt.Println("Разность: ", a-b)
	//	fmt.Println("Умножение: ", a*b)
	//	fmt.Println("Деление: ", a/b)

	//c := 256
	//g := int8(c)
	//fmt.Println("Value g:", g)
	//
	//a := 10
	//b := (1 + 2 + 3 - 2 - 2) * 10 * (100 % 90)
	//c := a * 10
	//d := c * a / 2
	//e := d % b
	////fmt.Println('\n')
	//fmt.Println(e)
	//a := 10.4
	//b := int(a)
	//fmt.Println(int8(int32(int64(b))))
	//a := 100000000000.4
	//b := int(a)
	//fmt.Println("Result of different task: ")
	//fmt.Println(int8(int32(int64(b))))

	var distance float64
	var time int

	_, err := fmt.Scan(&distance, &time)
	if err != nil {
		fmt.Println("Ошибка при чтении входных данных:", err)
		return
	}
	speed := distance / float64(time)
	fmt.Println(speed)

}
