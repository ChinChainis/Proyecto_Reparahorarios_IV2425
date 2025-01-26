package models

import "fmt"

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
	for contador := 1; contador < 6; contador++ {
		var contenidoDeUnDia []FranjaHoraria = AnalizaSuperposicionAsignaturaEnDia(horarioCompleto, contador)
		if SuperposicionConjuntoAsignaturas(contenidoDeUnDia) {
			return true, fmt.Errorf("superposición encontrada en día %v", contador)
		}
	}
	return false, nil
}

func haceHorario(asignaturasEntrada []FranjaHoraria, dia DiaSemana) ([]FranjaHoraria, error) {
	var horarioSalida []FranjaHoraria
	var horarioEntradaDia = AnalizaSuperposicionAsignaturaEnDia(asignaturasEntrada, int(dia))
	horarioSalida = append(horarioSalida, asignaturasEntrada[0])
	for i := 1; i < len(horarioEntradaDia); i++ {
		if !ExisteSolapamientoAsignaturas(horarioEntradaDia[i-1], horarioEntradaDia[i]) {
			horarioSalida = append(horarioSalida, horarioEntradaDia[i])
		}
	}
	if SuperposicionConjuntoAsignaturas(horarioSalida) {
		return nil, fmt.Errorf("horario incompleto")
	} else {
		return horarioSalida, nil
	}
}
