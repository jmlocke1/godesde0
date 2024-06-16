package modelos

type Hombre struct {
	Humano
}

func (h *Hombre) Sexo() string { return "Hombre" }
