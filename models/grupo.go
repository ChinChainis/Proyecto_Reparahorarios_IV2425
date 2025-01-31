package models

import (
	"fmt"
	"slices"
)

type Grupo struct {
	Nombre  string
	Horario *Horario
}

func NewGrupo(nombre string) (*Grupo, error) {

	horario, err := NewHorario()
	if err != nil {
		return nil, err
	}

	return &Grupo{

		Nombre:  nombre,
		Horario: horario,
	}, nil
}

func ExisteSolapamientoAsignaturas(asignaturaPrimera FranjaHoraria, asignaturaSecundaria FranjaHoraria) bool {
	if asignaturaPrimera.Inicio.Hora == asignaturaSecundaria.Inicio.Hora {
		return true
	} else if asignaturaPrimera.Inicio.Hora < asignaturaSecundaria.Inicio.Hora &&
		(asignaturaPrimera.Inicio.Hora+int(asignaturaPrimera.Duracion)) > asignaturaSecundaria.Inicio.Hora {
		return true
	} else if asignaturaPrimera.Inicio.Hora > asignaturaSecundaria.Inicio.Hora &&
		asignaturaPrimera.Inicio.Hora < (asignaturaSecundaria.Inicio.Hora+int(asignaturaSecundaria.Duracion)) {
		return true
	} else {
		return false
	}
}

func SuperposicionConjuntoAsignaturas(horariodeldia []FranjaHoraria) bool {
	for j := 0; j < len(horariodeldia)-1; j++ {
		var franjaActual FranjaHoraria = horariodeldia[j]
		var franjaSiguiente FranjaHoraria = horariodeldia[j+1]
		if ExisteSolapamientoAsignaturas(franjaActual, franjaSiguiente) {
			return true
		}
	}
	return false
}

func AnalizaSuperposicionAsignaturaEnDia(horarioEntrada []FranjaHoraria, dia int) []FranjaHoraria {
	var asignaturasDia []FranjaHoraria
	for i := 0; i < len(horarioEntrada); i++ {
		if int(horarioEntrada[i].Dia) == dia {
			asignaturasDia = append(asignaturasDia, horarioEntrada[i])
		}
	}
	return asignaturasDia
}

func CompruebaSuperposicionesEnSemana(horarioCompleto []FranjaHoraria) (bool, error) {
	diaLaboral := 1
	var estado bool
	for diaLaboral < 6 || estado {
		var contenidoDeUnDia []FranjaHoraria = AnalizaSuperposicionAsignaturaEnDia(horarioCompleto, diaLaboral)
		estado = SuperposicionConjuntoAsignaturas(contenidoDeUnDia)
	}
	if estado {
		return true, fmt.Errorf("superposición encontrada en día %v", diaLaboral)
	} else {
		return estado, nil
	}

}

func haceHorarioDeUnDia(asignaturasEntrada []FranjaHoraria, dia DiaSemana) ([]FranjaHoraria, error) {
	var horarioSalida []FranjaHoraria
	var horarioEntradaDia = AnalizaSuperposicionAsignaturaEnDia(asignaturasEntrada, int(dia))
	var superposicionEnDia bool
	superposicionEnDia, _ = CompruebaSuperposicionesEnSemana(horarioEntradaDia)
	if len(horarioEntradaDia) > 0 || superposicionEnDia {
		horarioSalida = horarioEntradaDia
	}

	if SuperposicionConjuntoAsignaturas(horarioSalida) {
		return nil, fmt.Errorf("horario con superposiciones en %v", dia)
	} else {
		return horarioSalida, nil
	}
}

func horarioDeUnaSemana(horarioCompleto []FranjaHoraria) ([]FranjaHoraria, error) {
	var contenidoSemana []FranjaHoraria
	var horarioSemanalResultante []FranjaHoraria
	for diaActual := 1; diaActual < 6; diaActual++ {
		contenidoSemana, _ = haceHorarioDeUnDia(horarioCompleto, DiaSemana(diaActual))
		horarioSemanalResultante = slices.Concat(horarioSemanalResultante, contenidoSemana)
	}
	var superposicionSemana, _ = CompruebaSuperposicionesEnSemana(contenidoSemana)
	if superposicionSemana {
		return nil, fmt.Errorf("horario con superposiciones")
	} else {
		return horarioSemanalResultante, nil
	}
}
