package models

import (
	"time"
	"errors"
)

type Horario struct {
	Dia        time.Weekday
	HoraInicio time.Time
	HoraFin    time.Time
	Grupo      Grupo
}


func newHorario(dia time.Weekday, horaInicio time.Time, horaFin time.Time, grupo Grupo) (*Horario, error) {

	
	if horaInicio.After(horaFin) {
		return nil, errors.New("la hora de inicio debe ser anterior a la hora de fin")
	}

	
	horario := &Horario{
		Dia:        dia,
		HoraInicio: horaInicio,
		HoraFin:    horaFin,
		Grupo:      grupo,
	}

	if err := ValidarHorarioLaboral(horario); err != nil {
		return nil, err
	}

	return horario, nil
}


func (h *Horario) ConflictoCon(other *Horario) bool {
	if h.Dia != other.Dia {
		return false
	}
	return h.HoraInicio.Before(other.HoraFin) && h.HoraFin.After(other.HoraInicio)
}
