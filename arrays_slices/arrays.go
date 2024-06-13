package arraysslices

import "fmt"

var tabla [10]int
var matriz [20][30]int

func MuestroArrays() {
	tabla[7] = 20
	tabla[2] = 567
	tabla[0] = 86
	fmt.Println(tabla)

	tabla2 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla2)
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	matriz[7][24] = 15
	fmt.Println(matriz)
}
