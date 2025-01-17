package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, err, tt.resultado, "Los resultados debería ser iguales")
	}

}
