package utiles

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	//"github.com/tealeg/xlsx"
	"github.com/extrame/xls"
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
		Addr:     "192.168.76.233:5432",
		User:     "dgevyl",
		Password: "qwerty",
		Database: "dgevyl",
	})
	defer db.Close()

	excelFileName := path
	/*xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		c.String(http.StatusOK, "error OpenFile : %s", err)
	}*/
	if xlFile, err := xls.Open(excelFileName, "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			fmt.Print("Total Lines ", sheet1.MaxRow, sheet1.Name)
			//col1 := sheet1.Row(0).Col(0)
			//col2 := sheet1.Row(0).Col(0)
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				if i > 0{
					/*var a [100]string
					var contador int = 0
					for _, cell := range row.Cells {
						a[contador] = cell.String()
						contador++
					}
					*/
					row1 := sheet1.Row(i)
					idcita, err := strconv.Atoi(row1.Col(0))
					if err != nil {
						c.String(http.StatusOK, "error convirtiendo entero : %s", err)
					}

					datosPersonales := strings.Split(row1.Col(2), ",")
					idtipodoc := 0
					switch row1.Col(3) {
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
					fechaCita := strings.Split(row1.Col(11), " ")
					fechaCitaEditada := strings.Split(fechaCita[0], "/")
					new_sigeci := &Sigeci{
						Idcita:    idcita,
						Idtipodoc: idtipodoc,
						Numdoc:    row1.Col(4),
						Nombre:    datosPersonales[1],
						Apellido:  datosPersonales[0],
						Fecha:     fechaCitaEditada[2]+"-"+fechaCitaEditada[1]+"-"+fechaCitaEditada[0],
						Hora:      fechaCita[1],
						Metadata:  `{"nacionalidad":"Argentina"}`,
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
			}
		}
	}
	/*
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
					Metadata:  `{"nacionalidad":"Argentina"}`,
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
	}*/
	c.String(http.StatusOK, "temino")

}
