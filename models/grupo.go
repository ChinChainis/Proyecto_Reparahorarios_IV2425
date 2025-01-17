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

func NoSuperposicionDeAsignaturasEnGrupo(horarioEntrada Horario) bool {
	var haySuperposicion bool = false
	for i := 0; i < len(horarioEntrada.Horario); i++ {
		var franjaActual FranjaHoraria = horarioEntrada.Horario[i]

		for j := i + 1; j < len(horarioEntrada.Horario); j++ {
			var franjaAComparar FranjaHoraria = horarioEntrada.Horario[j]
			if franjaActual.Dia == franjaAComparar.Dia {
				if franjaActual.Inicio.Hora == franjaAComparar.Inicio.Hora {
					haySuperposicion = true
				} else if franjaActual.Inicio.Hora < franjaAComparar.Inicio.Hora &&
					(franjaActual.Inicio.Hora+int(franjaActual.Duracion)) > franjaAComparar.Inicio.Hora {
					haySuperposicion = true
				} else if franjaActual.Inicio.Hora > franjaAComparar.Inicio.Hora &&
					franjaActual.Inicio.Hora < (franjaAComparar.Inicio.Hora+int(franjaAComparar.Duracion)) {
					haySuperposicion = true
				}
			}

		}

	}
	return haySuperposicion
}
