package modelos

type Mujer struct {
	Humano
}

func (m *Mujer) Sexo() string { return "Mujer" }
