package models

type Grupo struct {
	Nombre  string
	Horario *Horario
}

func NewGrupo(nombre string) (*Grupo, error) {

	horario, err := NewHorario()
	if err != nil {
		return nil, err
	}

	return &Grupo{

		Nombre:  nombre,
		Horario: horario,
		
	}, nil
}
