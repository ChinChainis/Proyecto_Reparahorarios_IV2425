package models

type Horario struct {
	Horario []FranjaHoraria
}

func NewHorario() (*Horario, error) {
	return &Horario{
		Horario: []FranjaHoraria{},
	}, nil
}
