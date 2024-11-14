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
		return nil, err
	}

	return estudiante, nil
}
