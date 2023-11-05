package forgotpassword

import (
	"encoding/json"
	"my-app/data/mysqlxdata"

	"github.com/labstack/echo/v4"
)

func SelectData(c echo.Context) error {
	// fmt.Println("masuk ke select data")
	// fmt.Println(c.Request().FormValue("email"))
	// fmt.Println(c.Request().FormValue("perusahaan"))
	db, err := mysqlxdata.ConnectOpenCareer()
	if err != nil {
		// fmt.Println("error nih: ")
		return nil
	}

	defer db.Close()

	rows, err := db.Query(`
	SELECT 
		id, 
		idregister, 
		email, 
		password, 
		DataPersonal 
		FROM tblapplicant 
	WHERE email = ? AND perusahaan = ?`, c.Request().FormValue("email"), c.Request().FormValue("perusahaan"))
	if err != nil {
		// fmt.Println("error nih: ", err)
		return nil
	}

	defer rows.Close()

	// fmt.Println(rows)

	var results map[string]interface{}
	for rows.Next() {
		var id int
		var idregister, email, password string
		var DataPersonal json.RawMessage

		err = rows.Scan(&id, &idregister, &email, &password, &DataPersonal)
		if err != nil {
			return nil
		}

		result := make(map[string]interface{})
		result["id"] = id
		result["idregister"] = idregister
		result["email"] = email
		result["DataPersonal"] = DataPersonal
		result["password"] = password

		// results = append(results, result)
		results = result
	}
	if results != nil {
		return c.JSON(200, results)
	}

	return nil
}
