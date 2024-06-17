package main

import (
	"fmt"

	"github.com/godesde0/goroutines"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoaTexto(1589734373446362525)
	// fmt.Println(estado)
	// fmt.Println(texto)
	/*
		if os := runtime.GOOS; os == "Linux." || os == "OS X." {
			fmt.Println("Esto no es Windows")
			fmt.Println(os)
		} else if os == "windows" {
			fmt.Println("Esto es Windows")
			fmt.Println(os)
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "windows":
			fmt.Println("Esto es Windows")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Println("El sistema operativo es", os)
		}

		numero, mensaje := ejercicios.CuentaString("+500")

		fmt.Println(numero)

		fmt.Println(mensaje)
		teclado.IngresoNumeros()
		iteraciones.Iterar()
		ejercicios.Multiplica()
	*/
	// ficheros.SumaTabla()
	// ficheros.LeoArchivo()
	// funciones.Calculos()
	// funciones.LlamarClosure()
	// funciones.Exponencia(2)
	/*
		for i := 11; i < 30; i++ {
			fmt.Println("Factorial de ", i, " : ", funciones.Factorial(i))
		}
	*/
	// arraysslices.Capacidad()
	// mapas.MostrarMapas()
	// users.AltaUsuario()
	// Pedro := new(modelos.Hombre)
	// e.HumanosRespirando(Pedro)
	// Marta := new(modelos.Mujer)
	// e.HumanosRespirando(Marta)
	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanic()
	canal1 := make(chan bool)
	canal2 := make(chan bool)
	go goroutines.MiNombreLentoooo("José Miguel", canal1)
	go goroutines.MiNombreLentoooo("Paquito Chocolatero", canal2)
	fmt.Println("Estoy aquí")
	<-canal1
	<-canal2
	// var x string
	// fmt.Scanln(&x)
}
