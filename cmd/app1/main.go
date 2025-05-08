package main

import (
	"fmt"
	"goProjects/internal/utils"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	str := "Hello, Go!"
	reversed := utils.Reverse(str)
	fmt.Println("Original:", str)
	fmt.Println("Reversed:", reversed)

}
