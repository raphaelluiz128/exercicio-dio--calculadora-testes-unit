package main

import (
	"testing"
)

func TestCalculadora(t *testing.T) {
	tests := []struct {
		num1     float64
		num2     float64
		operator string
		expected float64
		err      bool // Indica se um erro é esperado, como divisão por zero
	}{
		{10, 5, "+", 15, false},
		{10, 5, "-", 5, false},
		{10, 5, "*", 50, false},
		{10, 5, "/", 2, false},
		{10, 0, "/", 0, true}, // Divisão por zero
		{10, 5, "%", 0, true}, // Operador inválido
	}

	for _, test := range tests {
		t.Run(test.operator, func(t *testing.T) {
			var result float64
			var err bool

			switch test.operator {
			case "+":
				result = test.num1 + test.num2
			case "-":
				result = test.num1 - test.num2
			case "*":
				result = test.num1 * test.num2
			case "/":
				if test.num2 != 0 {
					result = test.num1 / test.num2
				} else {
					err = true
				}
			default:
				err = true
			}

			if err != test.err {
				t.Errorf("Erro esperado: %v, mas obteve: %v", test.err, err)
			}

			if !err && result != test.expected {
				t.Errorf("Esperado: %.2f, obteve: %.2f", test.expected, result)
			}
		})
	}
}
