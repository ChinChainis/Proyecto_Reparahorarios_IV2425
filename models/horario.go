package models

type Horario struct {
	Horario []FranjaHoraria
}

func NewHorario() (*Horario, error) {
	return &Horario{
		Horario: []FranjaHoraria{},
	}, nil
}

func AsignaturaRepetida(horario Horario) bool {
	var hayRepeticion bool = false
	for i := 0; i < len(horario.Horario); i++ {
		for j := i + 1; j < len(horario.Horario); j++ {
			if horario.Horario[i].Asignatura == horario.Horario[j].Asignatura && horario.Horario[i].Turno == horario.Horario[j].Turno {
				hayRepeticion = true
			}
		}
	}
	return hayRepeticion
}
