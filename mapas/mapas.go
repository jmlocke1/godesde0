package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D. F."
	paises["Argentina"] = "Buenos Aires"
	paises["España"] = "Madrid"
	fmt.Println(paises)
	fmt.Println(paises["España"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	campeonato["Zaragoza"] = 19
	fmt.Println(campeonato)
	// Recorrer un mapa con un for
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de %d\n", equipo, puntaje)
	}
	// Borrar un elemento de un map
	delete(campeonato, "Real madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t\n", puntaje, existe)
	puntaje, existe = campeonato["Zaragoza"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t\n", puntaje, existe)
	var s string
	fmt.Println(s)
}
