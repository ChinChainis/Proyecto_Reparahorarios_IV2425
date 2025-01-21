package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperposicionHorasAsignaturaCompletamenteSuperpuesta(t *testing.T) {
	asignaturasTesteo := []struct {
		franjas   []FranjaHoraria
		resultado bool
	}{
		{franjas: []FranjaHoraria{{Lunes, HoraMinuto{9, MinutoCero}, DosHoras, "fisica", Manana},
			{Lunes, HoraMinuto{10, MinutoCero}, UnaHora, "lengua", Manana}}, resultado: true},
	}
	for _, tt := range asignaturasTesteo {
		err := NoSuperposicionDeAsignaturasEnGrupo(tt.franjas)
		assert.Equal(t, err, tt.resultado, "Una asignatura ocupa el horario de otra completamente")
	}
}

func TestSuperposicionHorasAsignaturaParcialmenteSuperpuesta(t *testing.T) {
	asignaturasTesteo := []struct {
		franjas   []FranjaHoraria
		resultado bool
	}{
		{franjas: []FranjaHoraria{{Lunes, HoraMinuto{12, MinutoCero}, DosHoras, "fisica", Manana},
			{Lunes, HoraMinuto{11, MinutoCero}, DosHoras, "lengua", Manana}}, resultado: true},
	}
	for _, tt := range asignaturasTesteo {
		err := NoSuperposicionDeAsignaturasEnGrupo(tt.franjas)
		assert.Equal(t, err, tt.resultado, "Una asignatura coincide parcialmente con otra")
	}
}

func TestSuperposicionHorasAsignaturaDiferentesTurnos(t *testing.T) {
	asignaturasTesteo := []struct {
		franjas   []FranjaHoraria
		resultado bool
	}{
		{franjas: []FranjaHoraria{{Lunes, HoraMinuto{12, MinutoCero}, DosHoras, "fisica", Manana},
			{Lunes, HoraMinuto{16, MinutoCero}, DosHoras, "lengua", Tarde}}, resultado: false},
	}
	for _, tt := range asignaturasTesteo {
		err := NoSuperposicionDeAsignaturasEnGrupo(tt.franjas)
		assert.Equal(t, err, tt.resultado, "Horario funciona correctamente")
	}
}
