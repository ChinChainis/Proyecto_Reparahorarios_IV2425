package models

import "errors"

type Estudiante struct {
	Nombre      string
	Asignaturas []Asignatura
}

func newEstudiante(nombre string, asignaturas []Asignatura) *Estudiante {
	return &Estudiante{
		Nombre:      nombre,
		Asignaturas: asignaturas,
	}
}