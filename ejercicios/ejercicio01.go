package ejercicios

import (
	"strconv"
)

func CuentaString(cadena string) (int, string) {
	numero, error := strconv.Atoi(cadena)
	var resultado string
	if error == nil {
		if numero > 100 {
			resultado = "Es mayor que 100"
		} else {
			resultado = "Es menor o igual que 100"
		}
	} else {
		resultado = "Ha ocurrido un error: " + error.Error()
	}

	return numero, resultado
}
