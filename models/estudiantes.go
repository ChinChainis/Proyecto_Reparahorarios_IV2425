package models

import (
	"errors"
)

type Estudiante struct {
	Nombre      string
	Asignaturas []Asignatura
}


func newEstudiante(nombre string, asignaturas []Asignatura) (*Estudiante, error) {
	estudiante := &Estudiante{
		Nombre:      nombre,
		Asignaturas: asignaturas,
	}

	if err := ValidarConflictosDeHorario(estudiante); err != nil {
		return nil, fmt.Errorf("error al validar conflictos de horario para el estudiante %q: %w", nombre, err)
	}

	return estudiante, nil
}
