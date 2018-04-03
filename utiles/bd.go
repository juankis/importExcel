package utiles

import (
	//"fmt"
	//"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	//"github.com/go-pg/pg/orm"
)

func ExampleDB_Model() {
	db := pg.Connect(&pg.Options{
		Addr:     "192.168.76.233:5432",
		User:     "dgevyl",
		Password: "qwerty",
		Database: "dgevyl",
	})
	defer db.Close()
}
	/*
		err := createSchema(db, c)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error creando la conexion: %s", err))
			//panic(err)
		}
	*/
	/*
	user1 := &User{
		Name:   "admin",
		Emails: "admin2@admin",
	}

	err := db.Insert(user1)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error insertando: %s", err))
		//panic(err)
	}

	// Select user by primary key.
	user := &User{Id: user1.Id}
	err = db.Select(user)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error selected : %s", err))
		//panic(err)
	}

	// Select all users.
	var users []User
	err = db.Model(&users).Select()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error seleccionando usuarios : %s", err))
		//panic(err)
	}

	fmt.Println(user)
	fmt.Println(users)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}

func createSchema(db *pg.DB, c *gin.Context) error {
	for _, model := range []interface{}{(*Turno)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error createSchema : %s", err))
		}
	}
	return nil
}
*/
