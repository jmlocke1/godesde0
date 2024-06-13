package funciones

import (
	"fmt"
)

func Exponencia(valor int) {
	if valor > 100000000 {
		return
	}
	fmt.Println(valor)
	Exponencia(valor * 2)
}

func Factorial(n int) int {
	if n == 1 {
		return n
	} else {
		return n * Factorial(n-1)
	}
}
