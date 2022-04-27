package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go Calculator!")
	cmd := readLine("Enter command: [a]dd, [s]ubstract, [m]ultiply, [d]ivide: ")
	fmt.Println(cmd)
	if cmd == "a" {
		num1, num2 := getUserNumbers()
		result := num1 + num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "s" {
		num1, num2 := getUserNumbers()
		result := num1 - num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "m" {
		num1, num2 := getUserNumbers()
		result := num1 * num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "d" {
		num1, num2 := getUserNumbers()
		result := float32(num1) / float32(num2)
		sResult := fmt.Sprintf("%f", result)
		fmt.Println(sResult)
	} else {
		fmt.Println("Invalid input ;)")
	}
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	num1string := readLine("First number: ")
	num1, err := strconv.Atoi(num1string)
	if err != nil {
		fmt.Println(err)
	}
	num2string := readLine("Second number: ")
	num2, err := strconv.Atoi(num2string)
	if err != nil {
		fmt.Println(err)
	}
	return num1, num2
}
