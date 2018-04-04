package utiles

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/tealeg/xlsx"
)

func moverArchivo(c *gin.Context) string {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	}
	nameFile := strconv.Itoa(rand.Intn(10000)) + "_" + file.Filename
	if err := c.SaveUploadedFile(file, nameFile); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	}
	c.String(http.StatusOK, fmt.Sprintf("El archivo %s ha sido trasladado con exito", file.Filename))
	return nameFile
}

func ImportExcel(c *gin.Context) {
	excelFileName := moverArchivo(c)
	db := pg.Connect(&pg.Options{
		Addr:     "192.168.76.233:5432",
		User:     "dgevyl",
		Password: "qwerty",
		Database: "dgevyl",
	})
	defer db.Close()

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		c.String(http.StatusOK, "error OpenFile : %s", err)
	}

	for _, sheet := range xlFile.Sheets {
		//var row_cont int = 0
		for _, row := range sheet.Rows {
			//if row_cont > 0 {
				var a [100]string
				var contador int = 0
				for _, cell := range row.Cells {
					a[contador] = cell.String()
					contador++
				}
				new_sigeci := getNewSigeci(a)

				sigeci := &Sigeci{Idcita: new_sigeci.Idcita}
				err = db.Select(sigeci)
		     if err != nil {
		             c.String(http.StatusBadRequest, fmt.Sprintf("Error selected : %s Idcita: %v", err, sigeci.Idcita))
		     }
				 if sigeci.Fecha != "" {
		             c.String(http.StatusBadRequest, fmt.Sprintf("El turno ya existe : %s", sigeci.Idcita))
		     } else {
		             err := db.Insert(new_sigeci)
		             if err != nil {
		                     c.String(http.StatusBadRequest, fmt.Sprintf("Error insertando: %s", err))
		             } else {
		                     c.String(http.StatusBadRequest, fmt.Sprintf("insertado"))
		             }
		     }
			}
			//row_cont++}
	}
	c.String(http.StatusOK, "temino")
}

func getTipoDoc(tipo_doc string) int{
	idtipodoc := 0

	switch tipo_doc {
	case "DNI":
		idtipodoc = 1
	case "DNI_EXT":
		idtipodoc = 5
	case "PASAPORTE":
		idtipodoc = 6
	case "LE":
		idtipodoc = 2
	case "LC":
		idtipodoc = 3
	case "CI":
		idtipodoc = 4
	}

	return idtipodoc
}

func getNewSigeci(a [100]string) *Sigeci{
	idcita, err := strconv.Atoi(a[0])
	if err != nil {
		//c.String(http.StatusOK, "error convirtiendo entero : %s", err)
		panic("error idcita getNewSigeci()")
	}
	datosPersonales := strings.Split(a[2], ",")
	idtipodoc := getTipoDoc(a[3])

	fechaCita := strings.Split(a[11], " ")
	fechaCitaEditada := strings.Split(fechaCita[0], "/")
	new_sigeci := &Sigeci{
		Idcita:    idcita,
		Idtipodoc: idtipodoc,
		Numdoc:    a[4],
		Nombre:    datosPersonales[1],
		Apellido:  datosPersonales[0],
		Fecha:     fechaCitaEditada[2]+"-"+fechaCitaEditada[1]+"-"+fechaCitaEditada[0],
		Hora:      fechaCita[1],
		Metadata:  `{"nacionalidad":"`+a[15]+`""}`,
	}
	return new_sigeci
}
