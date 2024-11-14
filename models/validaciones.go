package models

import "errors"

// ValidarHorarios verifica que no haya solapamientos entre los horarios de las asignaturas de un estudiante.
func ValidarHorarios(e *Estudiante) error {
	for i := 0; i < len(e.Asignaturas); i++ {
		for j := i + 1; j < len(e.Asignaturas); j++ {
			if err := e.Asignaturas[i].ValidarSolapamiento(&e.Asignaturas[j]); err != nil {
				return err
			}
		}
	}
	return nil
}

// ValidarSolapamiento verifica que no haya conflicto entre los horarios de esta asignatura y otra asignatura.
func (a *Asignatura) ValidarSolapamiento(other *Asignatura) error {
	for _, h1 := range a.Horarios {
		for _, h2 := range other.Horarios {
			if h1.ConflictoCon(&h2) {
				return errors.New("conflicto de horario entre " + a.Nombre + " y " + other.Nombre)
			}
		}
	}
	return nil
}
