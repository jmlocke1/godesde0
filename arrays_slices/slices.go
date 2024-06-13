package arraysslices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arr [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arr[3:]   // Slice creado desde un vector, desde la posición 3
	porcion2 := arr[:5]  // Slice creado desde un vector, desde la posición 0 hasta la 5
	porcion3 := arr[4:7] // Slice creado desde un vector, desde la posición 4 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
		fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
	}
}
