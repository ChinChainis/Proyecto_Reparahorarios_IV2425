package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperposicionHoras(t *testing.T) {
	asignaturasTesteo := []struct {
		franjas   []FranjaHoraria
		resultado bool
	}{
		{franjas: []FranjaHoraria{{Lunes, HoraMinuto{8, MinutoCero}, DosHoras, "fisica", Manana},
			{Lunes, HoraMinuto{12, MinutoCero}, DosHoras, "mates", Manana},
			{Martes, HoraMinuto{8, MinutoCero}, DosHoras, "lengua", Manana}}, resultado: false},
	}
	for _, tt := range asignaturasTesteo {
		res, err := CompruebaSuperposicionesEnSemana(tt.franjas)
		assert.Equal(t, tt.resultado, res, "Horario funciona correctamente")
		assert.NoError(t, err)
	}
}

func TestGeneracionHorario(t *testing.T) {
	franjasTesteo := []struct {
		franjas []FranjaHoraria
		dia     DiaSemana
	}{
		{franjas: []FranjaHoraria{{Lunes, HoraMinuto{8, MinutoCero}, DosHoras, "fisica", Manana},
			{Lunes, HoraMinuto{12, MinutoCero}, DosHoras, "mates", Manana},
			{Martes, HoraMinuto{8, MinutoCero}, DosHoras, "lengua", Manana},
			{Martes, HoraMinuto{10, MinutoCero}, DosHoras, "mates", Manana},
			{Martes, HoraMinuto{16, MinutoCero}, UnaHora, "fisica", Tarde}}, dia: Martes},
	}
	for _, tt := range franjasTesteo {
		res, err := haceHorario(tt.franjas, tt.dia)
		assert.Equal(t, len(res), 3, "horario esperado")
		assert.NoError(t, err)
	}
}
