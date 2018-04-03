package utiles

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func ImportExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	nameFile := file.Filename + "_" + strconv.Itoa(rand.Intn(10000))
	if err := c.SaveUploadedFile(file, nameFile); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("El archivo %s ha sido importados con exito", file.Filename))
}

func GuardarFilas(path string, c *gin.Context) {
	excelFileName := path
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		c.String(http.StatusOK, "error: %s", err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var a [6]string
			contador := 0
			for _, cell := range row.Cells {
				a[contador] = cell.String()
				contador++
			}
			guardarCita(a)
		}
	}
	c.String(http.StatusOK, "temino")
}

func guardarCita(a [6]string) {
	i, err := strconv.Atoi(a[0])
	turno := &Turno{
		Idcita:       1,
		Idtipodoc:    "admin2@admin",
		Numdoc:       "admin2@admin",
		Nombre:       "admin2@admin",
		Apellido:     "admin2@admin",
		Nacionalidad: "admin2@admin",
	}
}
