package models

import "errors"

type Turno string

const (
	Maniana Turno = "Maniana"
	Tarde  Turno = "Tarde"
)

type Grupo struct {
	Nombre string
	Turno  Turno
}

func NewGrupo(nombre string, turno Turno) (*Grupo, error) {
	// Validamos que el turno sea "Mañana" o "Tarde"
	if turno != Maniana && turno != Tarde {
		return nil, errors.New("turno inválido: debe ser 'Maniana' o 'Tarde'")
	}

	return &Grupo{
		Nombre: nombre,
		Turno:  turno,
	}, nil
}
