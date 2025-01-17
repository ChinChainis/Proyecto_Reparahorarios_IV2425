package models

import (
	"testing"
)

func TestCorrespondenciaHoraTurno(t *testing.T) {
	horaTesteo := []struct {
		hora      int
		turno     Turno
		resultado bool
	}{
		{11, Manana, true},
		{17, Tarde, true},
		{15, Manana, false},
		{8, Tarde, false},
	}

	for _, tt := range horaTesteo {
		err := CompruebaCorrespondenciaHoraTurno(
			tt.hora,
			tt.turno,
		)
		if resultado := err; resultado != tt.resultado {
			t.Errorf("esHoraValidaPorTurno(%v, %v) = %v; resultado esperado: %v", tt.hora, tt.turno, resultado, tt.resultado)
		}
	}

}
