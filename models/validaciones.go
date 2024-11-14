
package models

import (
	"errors"
	"fmt"
	"time"
)


func ValidarConflictosDeHorario(estudiante *Estudiante) error {
	for i := 0; i < len(estudiante.Asignaturas); i++ {
		for j := 0; j < len(estudiante.Asignaturas[i].Horarios); j++ {
			horarioActual := estudiante.Asignaturas[i].Horarios[j]
			
			
			for k := i; k < len(estudiante.Asignaturas); k++ {
				for l := 0; l < len(estudiante.Asignaturas[k].Horarios); l++ {
					horarioComparar := estudiante.Asignaturas[k].Horarios[l]
					
			
					if i == k && j == l {
						continue
					}

					
					if horarioActual.ConflictoCon(&horarioComparar) {
						return fmt.Errorf("conflicto de horario: %s en %s tiene un solapamiento con %s en %s",
							estudiante.Asignaturas[i].Nombre, horarioActual.Grupo.Nombre,
							estudiante.Asignaturas[k].Nombre, horarioComparar.Grupo.Nombre)
					}
				}
			}
		}
	}
	return nil
}



func ValidarHorarioLaboral(horario *Horario) error {
	if horario.Dia == time.Saturday || horario.Dia == time.Sunday {
		return errors.New("el horario debe ser de lunes a viernes")
	}


	manianaInicio := time.Date(0, 0, 0, 8, 30, 0, 0, time.UTC)
	manianaFin := time.Date(0, 0, 0, 14, 30, 0, 0, time.UTC)
	tardeInicio := time.Date(0, 0, 0, 15, 30, 0, 0, time.UTC)
	tardeFin := time.Date(0, 0, 0, 20, 30, 0, 0, time.UTC)


	if horario.Grupo.Turno == Maniana {
		if horario.HoraInicio.Before(manianaInicio) || horario.HoraFin.After(manianaFin) {
			return errors.New("el horario de mañana debe ser entre las 08:30 y las 14:30")
		}
	}

	if horario.Grupo.Turno == Tarde {
		if horario.HoraInicio.Before(tardeInicio) || horario.HoraFin.After(tardeFin) {
			return errors.New("el horario de tarde debe ser entre las 15:30 y las 20:30")
		}
	}

	return nil
}
