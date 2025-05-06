package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

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

	// обратите внимание на одинарные кавычки
	var symbol rune = 'A'
	fmt.Println(symbol)

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

}
