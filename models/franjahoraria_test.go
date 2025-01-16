package models

import (
	"testing"
)

func TestEsHoraValidaPorTurno(t *testing.T) {
	type horaTesteo struct {
		hora      int
		turno     Turno
		resultado bool
	}
	var prueba horaTesteo = horaTesteo{11, "maniana", true}

	if resultado := EsHoraValidaPorTurno(prueba.hora, prueba.turno); resultado != prueba.resultado {
		t.Errorf("esHoraValidaPorTurno(%v, %v) = %v; resultado esperado: %v", prueba.hora, prueba.turno, resultado, prueba.resultado)
	}

}
