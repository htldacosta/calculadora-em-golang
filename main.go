package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fmt.Println("Enter the operation (using this format: 2*2):")
	var input string
	fmt.Scan(&input)

	operation := strings.Split(input, "")
	fmt.Println(operation[0])
	fmt.Println(operation[1])
	fmt.Println(operation[2])
	
}

func getResult(operation []String) int {

	switch operation[1] {
	case "+":
		return operation[0] + operation[2]
	}
}