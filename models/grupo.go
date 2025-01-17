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

	if NoSuperposicionDeAsignaturasEnGrupo(horario.Horario) {
		return nil, fmt.Errorf("encontradas asignaturas que se superponen")
	}

	return &Grupo{

		Nombre:  nombre,
		Horario: horario,
	}, nil
}

func NoSuperposicionDeAsignaturasEnGrupo(horarioEntrada []FranjaHoraria) bool {
	var haySuperposicion bool = false
	for i := 0; i < len(horarioEntrada); i++ {
		var franjaActual FranjaHoraria = horarioEntrada[i]

		for j := i + 1; j < len(horarioEntrada); j++ {
			var franjaAComparar FranjaHoraria = horarioEntrada[j]
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
