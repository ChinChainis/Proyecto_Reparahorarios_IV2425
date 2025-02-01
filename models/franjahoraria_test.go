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
		assert.Equal(t, err, tt.resultado, "Los resultados deberían ser iguales")
	}

}

func TestDuracionTurnoCorrecto(t *testing.T) {
	horaTesteo := []struct {
		hora  int
		turno Turno
	}{
		{0, Manana},
		{1, Manana},
		{2, Manana},
		{3, Manana},
		{4, Manana},
		{5, Manana},
		{6, Manana},
		{7, Manana},
		{14, Manana},
		{14, Tarde},
		{21, Tarde},
		{22, Tarde},
		{23, Tarde},
	}

	for _, tt := range horaTesteo {
		err := CompruebaCorrespondenciaHoraTurno(
			tt.hora,
			tt.turno,
		)
		assert.Equal(t, err, false, "Los resultados deben ser falsos")
	}

}
