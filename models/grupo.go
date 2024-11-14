package models

type Grupo struct {
	Nombre string
}

func newGrupo(nombre string) *Grupo {
	return &Grupo{
		Nombre: nombre,
	}
}
