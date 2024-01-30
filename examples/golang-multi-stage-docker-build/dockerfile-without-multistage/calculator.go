package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	fmt.Println("Calculator app...")
	for {
		// Read input from user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter any calculation (Ex: 1 + 2 (or) 2 * 5 -> maintain Spaces):")
		text, _ := reader.ReadString('\n')
		
		// Trim newline character from input
		text = strings.TrimSpace(text)
		
		// Check if user enters exit to quit the program
		if text =="exit" {
			break
		}
		
		// split left operand and right operand
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Try again")
			continue
		}
		
		// Convert operands to integers
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}
		
		// Compute using the operator as a basis
		var result int
		switch parts[1] {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right 
		case "/":
			result = left / right
		default:
			fmt.Println("Invalid operator. Try again.")
			continue
		}
		
		// Print result
		fmt.Printf("Result:%d\n", result)
	}
}	
						
