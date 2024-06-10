package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1589734373446362525)
	fmt.Println(estado)
	fmt.Println(texto)
}
