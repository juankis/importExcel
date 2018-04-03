package utiles

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"github.com/go-pg/pg"
)

func moverArchivo(c *gin.Context) string{
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
	path:= moverArchivo(c)
	c.String(http.StatusOK, "path: %s",path)
	db := pg.Connect(&pg.Options{
		Addr:     "192.168.76.233:5432",
		User:     "dgevyl",
		Password: "qwerty",
		Database: "dgevyl",
	})
	defer db.Close()

	//excelFileName := "/home/jojeda/go/src/github.com/juankis/importExcel/reporte_test.xls"
	//excelFileName := "/home/jojeda/Descargas/testfile.xlsx"
	excelFileName := path
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		c.String(http.StatusOK, "error OpenFile : %s", err)
	}

	for _, sheet := range xlFile.Sheets {
		var row_cont int = 0

		for _, row := range sheet.Rows {
			c.String(http.StatusOK, "ROW : %s", strconv.Itoa(row_cont))
			//if row_cont > 0 {
				var a [100]string
				var contador int = 0
				for _, cell := range row.Cells {
					a[contador] = cell.String()
					contador++
				}

				entero,err := strconv.Atoi(a[0])
				//entero := uint64(entero)
				new_sigeci := &Sigeci{
					Idcita:       entero,
					Idtipodoc:    1,
					Numdoc:       a[4],
					Nombre:       a[2],
					Apellido:     a[2],
					Fecha:				"2018-02-02",
					Hora:					"10:00:00",
				}

				sigeci := &Sigeci{Idcita: new_sigeci.Idcita}
				err = db.Select(sigeci)
				if err != nil {
					c.String(http.StatusBadRequest, fmt.Sprintf("Error selected : %s Idcita: %v", err, sigeci.Idcita))
					//panic(err)
				}
				//numero, err := strconv.ParseInt(sigeci.Idcita, 64, 64)
				panic(sigeci)
				if sigeci != nil{
					c.String(http.StatusBadRequest, fmt.Sprintf("El turno ya existe : %s", sigeci.Idcita))
				}else{
					err := db.Insert(new_sigeci)
					if err != nil {
						c.String(http.StatusBadRequest, fmt.Sprintf("Error insertando: %s", err))
					}else{
						c.String(http.StatusBadRequest, fmt.Sprintf("insertado"))
					}
				}
			//}
			row_cont++
		}
	}
	c.String(http.StatusOK, "temino")

}
