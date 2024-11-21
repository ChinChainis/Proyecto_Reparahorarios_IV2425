package models

type Horario struct {
    Horario []FranjaHoraria
}

func NewHorario() *Horario {
    return &Horario{
        Horario: []FranjaHoraria{},
    }
}
