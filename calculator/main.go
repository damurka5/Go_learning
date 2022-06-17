package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin) // reader with 4096 bytes buffer size
	fmt.Println("Simple calculator")
	fmt.Println("Supports basic arithmetic operations: +, -, *, /, ln, sin, cos, tan (arguments are in degrees)")
	fmt.Println("Please, insert an expression")

	for { // infinite loop for infinite reading
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		inputArr := strings.Fields(text) // converting string an array of strings
		if strings.Compare(inputArr[0], "ln") == 0 {
			number1, _ := strconv.ParseFloat(inputArr[1], 32)
			fmt.Println(math.Log(number1))
		} else if strings.Compare(inputArr[0], "sin") == 0 {
			number1, _ := strconv.ParseFloat(inputArr[1], 32)
			fmt.Println(math.Sin(number1 * math.Pi / 180))
		} else if strings.Compare(inputArr[0], "cos") == 0 {
			number1, _ := strconv.ParseFloat(inputArr[1], 32)
			fmt.Println(math.Cos(number1 * math.Pi / 180))
		} else if strings.Compare(inputArr[0], "tan") == 0 {
			number1, _ := strconv.ParseFloat(inputArr[1], 32)
			fmt.Println(math.Tan(number1 * math.Pi / 180))
		} else {
			number1, _ := strconv.ParseFloat(inputArr[0], 32)
			number2, _ := strconv.ParseFloat(inputArr[0], 32)
			switch inputArr[1] {
			case "+":
				fmt.Println(number1 + number2)
			case "-":
				fmt.Println(number1 - number2)
			case "*":
				fmt.Println(number1 * number2)
			case "/":
				fmt.Println(number1 / number2)
			case "^":
				fmt.Println(math.Pow(number1, number2))
			default:
				fmt.Println("Unsupported operation")
			}
		}
	}
}
