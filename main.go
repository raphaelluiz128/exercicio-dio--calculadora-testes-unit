package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Bem-vindo à calculadora!")
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Resultado: %.2f\n", num1/num2)
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida.")
		}
	default:
		fmt.Println("Operador inválido! Use +, -, * ou /.")
	}
}
