package ficheros

import (
	"bufio"
	"fmt"
	"os"

	"github.com/godesde0/ejercicios"
)

var fileName string = "./ficheros/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Multiplica()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Multiplica()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error durante el Append " + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivoBad() {
	archivo, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
