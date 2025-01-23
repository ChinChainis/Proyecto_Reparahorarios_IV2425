package models

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

func ComparaSolapamientoAsignaturas(asignaturaPrimera FranjaHoraria, asignaturaSecundaria FranjaHoraria) bool {
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
		if ComparaSolapamientoAsignaturas(franjaActual, franjaSiguiente) {
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
