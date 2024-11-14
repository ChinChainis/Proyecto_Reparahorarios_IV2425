package models

type Grupo struct {
	Nombre string
	Turno string
}

func newGrupo(nombre, turno string) *Grupo {
	return &Grupo{
		Nombre: nombre,
		Turno: turno,
	}
}
