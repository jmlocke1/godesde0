package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplica() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var texto string
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
		texto += fmt.Sprintln(numero, " x ", i, "=", i*numero)
	}
	texto += fmt.Sprintln("") // Insertamos una línea de separación
	return texto
}
