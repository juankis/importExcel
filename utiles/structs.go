package utiles

import "fmt"

type Sigeci struct {
	tableName          struct{} `sql:"sigeci"`
	Idcita       int `sql:",pk"`
	Idtipodoc    int32
	Numdoc       string
	Nombre       string
	Apellido     string
	Fecha 			 string
	Hora				 string
}

func (t Sigeci) String() string {
	return fmt.Sprintf("Sigeci<%d %s %v %d %s %s>", t.Idcita, t.Idtipodoc, t.Numdoc, t.Nombre, t.Apellido)
}
