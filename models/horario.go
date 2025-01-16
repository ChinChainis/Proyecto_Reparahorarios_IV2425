package models

import "fmt"

type Horario struct {
	Horario []FranjaHoraria
}

func NewHorario(horario []FranjaHoraria) (*Horario, error) {
	if !AsignaturaRepetida(horario) {
		return nil, fmt.Errorf("Error de lógica interna de las asignaturas")
	}
	return &Horario{
		Horario: horario,
	}, nil
}

func AsignaturaRepetida(horario []FranjaHoraria) bool {
	var noHayRepeticion bool = true
	for i := 0; i < len(horario); i++ {
		for j := i + 1; j < len(horario); j++ {
			if horario[i].Asignatura == horario[j].Asignatura && horario[i].Turno == horario[j].Turno {
				noHayRepeticion = false
			}
		}
	}
	return noHayRepeticion
}
