package ejer_interfaces

import (
	"fmt"

	"github.com/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	fmt.Printf("Soy %s %s, y estoy respirando %t\n", articulo(hu), hu.Sexo(), hu.Respirar())
}

func articulo(hu interfaces.Humano) string {
	var articulo string
	if hu.Sexo() == "Hombre" {
		articulo = "un"
	} else {
		articulo = "una"
	}
	return articulo
}
