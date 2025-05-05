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
	//
	//fmt.Println("*")
	//fmt.Println("**")
	//fmt.Println("***")
	//fmt.Println("****")
	//fmt.Println("*****")
	//fmt.Println("******")
	//fmt.Println("*******")

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
}
