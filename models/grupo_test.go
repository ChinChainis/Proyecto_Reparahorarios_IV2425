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
		assert.Equal(t, err, tt.resultado, "Los resultados deberían ser iguales")
	}
}
