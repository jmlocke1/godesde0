package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplica() {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	fmt.Println("Introduzca un número para calcular su tabla de multiplicar")
	for {
		if scanner.Scan() {
			numtemp, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El número ingresado es incorrecto, introduzca otro número")
			} else {
				numero = numtemp
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(numero, " x ", i, "=", i*numero)
	}
}
