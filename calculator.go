package main

import(
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Go Calculator!")
	cmd := readLine("Enter any command: [+], [*], [-], [/]: ")
	fmt.Print(cmd)

	if cmd == "+" {
		num1, num2 := getUserNumbers()
		result := num1 + num2
		fmt.Println(fmt.Sprintf("%d", result))
	} else if cmd == "-" {
		num1, num2 := getUserNumbers()
		result := num1 - num2
		fmt.Println(fmt.Sprintf("%d", result))
	} else if cmd == "*" {
		num1, num2 := getUserNumbers()
		result := num1 * num2
		fmt.Println(fmt.Sprintf("%d", result))
	} else if cmd == "/" {
		num1, num2 := getUserNumbers()
		result := float32(num1) / float32(num2)
		fmt.Println(fmt.Sprintf("%f", result))
	} else {
		fmt.Println("Invalid input")
	}
	readLine("press 'enter' to calculat another time")
}

func readLine(message string) string{
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	num1String := readLine("First Number: ")
	num1, _ := strconv.Atoi(num1String) //strconv.Atoi => to convert int to string
	num2String := readLine("Second Number: ")
	num2, _ := strconv.Atoi(num2String)
	return num1, num2
}