package utiles

import "fmt"

type Turno struct {
	Idcita       int64
	Idtipodoc    string
	Numdoc       string
	Nombre       string
	Apellido     string
	Nacionalidad string
}

func (t Turno) String() string {
	return fmt.Sprintf("Turno<%d %s %v %d %s %s>", t.Idcita, t.Idtipodoc, t.Numdoc, t.Nombre, t.Apellido, t.Nacionalidad)
}
