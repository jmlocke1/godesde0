package modelos

type Humano struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Humano) Respirar() bool {
	h.respirando = true
	return h.respirando
}
func (h *Humano) Comer()  { h.comiendo = true }
func (h *Humano) Pensar() { h.pensando = true }

func (h *Humano) EstaVivo() bool {
	h.vivo = true
	return h.vivo
}
