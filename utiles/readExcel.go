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
	path := moverArchivo(c)
	c.String(http.StatusOK, "path: %s", path)
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "root",
		Database: "example",
	})
	defer db.Close()

	excelFileName := path
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		c.String(http.StatusOK, "error OpenFile : %s", err)
	}

	for _, sheet := range xlFile.Sheets {
		var row_cont int = 0
		for _, row := range sheet.Rows {
			c.String(http.StatusOK, "ROW : %s", strconv.Itoa(row_cont))
			if row_cont > 0 {
				var a [100]string
				var contador int = 0
				for _, cell := range row.Cells {
					a[contador] = cell.String()
					contador++
				}

				idcita, err := strconv.Atoi(a[0])
				if err != nil {
					c.String(http.StatusOK, "error convirtiendo entero : %s", err)
				}

				datosPersonales := strings.Split(a[2], ",")
				idtipodoc := 0
				switch a[3] {
				case "DNI":
					idtipodoc = 1
				case "DNI_EXT":
					idtipodoc = 2
				case "PASAPORTE":
					idtipodoc = 3
				}
				fechaCita := strings.Split(a[11], " ")

				new_sigeci := &Sigeci{
					Idcita:    idcita,
					Idtipodoc: idtipodoc,
					Numdoc:    a[4],
					Nombre:    datosPersonales[1],
					Apellido:  datosPersonales[0],
					Fecha:     fechaCita[0],
					Hora:      fechaCita[1],
				}

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
			row_cont++
		}
	}
	c.String(http.StatusOK, "temino")

}
