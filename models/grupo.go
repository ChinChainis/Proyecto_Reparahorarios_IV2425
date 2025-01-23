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
