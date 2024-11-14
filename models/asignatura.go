package models

type Asignatura struct {
	Nombre string
	Creditos int
	Horarios []Horario
}

func newAsignatura(nombre string, creditos int, horarios []Horario) *Asignatura {
	return &Asignatura{
		Nombre: nombre,
		Creditos: creditos,
		Horarios: horarios,
	}
}

