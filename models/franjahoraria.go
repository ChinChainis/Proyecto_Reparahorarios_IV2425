package models

import (
	"fmt"
)

type Minuto int

const (
	MinutoCero    Minuto = 0
	MinutoTreinta Minuto = 30
)

type Hora int

const (
	EntradaManana Hora = 8
	SalidaManana  Hora = 14
	EntradaTarde  Hora = 15
	SalidaTarde   Hora = 20
)

type DiaSemana int

const (
	Lunes DiaSemana = iota + 1
	Martes
	Miercoles
	Jueves
	Viernes
)

type HoraMinuto struct {
	Hora   int
	Minuto Minuto
}

type Turno string

const (
	Manana Turno = "maniana"
	Tarde  Turno = "tarde"
)

type Duracion int

const (
	UnaHora  Duracion = 1
	DosHoras Duracion = 2
)

type FranjaHoraria struct {
	Dia        DiaSemana
	Inicio     HoraMinuto
	Duracion   Duracion
	Asignatura string
	Turno      Turno
}

func NewFranjaHoraria(dia DiaSemana, inicio HoraMinuto, duracion Duracion, asignatura string, turno Turno) (*FranjaHoraria, error) {

	if !CompruebaCorrespondenciaHoraTurno(inicio.Hora, turno) {
		return nil, fmt.Errorf("hora inválida para el turno %s: %02d no está en el rango permitido", turno, inicio.Hora)
	}

	return &FranjaHoraria{
		Dia:        dia,
		Inicio:     inicio,
		Duracion:   duracion,
		Asignatura: asignatura,
		Turno:      turno,
	}, nil
}

func CompruebaCorrespondenciaHoraTurno(hora int, turno Turno) bool {
	switch turno {
	case Manana:
		return hora >= int(EntradaManana) && hora < int(SalidaManana)
	case Tarde:
		return hora >= int(EntradaTarde) && hora <= int(SalidaTarde)
	default:
		return false
	}
}
