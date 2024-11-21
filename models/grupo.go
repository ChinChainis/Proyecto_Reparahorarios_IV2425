package models

type Grupo struct {
	Nombre  string
	Horario Horario
}

func NewGrupo(nombre string) (*Grupo, error) {

	return &Grupo{
		Nombre:  nombre,
		Horario: *NewHorario(),
	}, nil
}
