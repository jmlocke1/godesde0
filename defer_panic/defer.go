package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Éste es un primer mensaje")
	defer fmt.Println("Éste es el mensaje final")
	fmt.Println("Éste es el segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic\n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontró el valor 1")
	}
}
