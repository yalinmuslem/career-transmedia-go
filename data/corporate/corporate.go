package corporate

import (
	"my-app/data/mysqlxdata"

	"github.com/labstack/echo/v4"
)

func SelectData(c echo.Context) error {
	db, err := mysqlxdata.ConnectOpenCareer()
	if err != nil {
		// fmt.Println("error nih: ")
		return nil
	}

	defer db.Close()

	rows, err := db.Query(`
	SELECT
		distinct(perusahaan) as perusahaan
	FROM tblapplicant
	GROUP BY perusahaan`)
	if err != nil {
		// fmt.Println("error nih: ", err)
		return nil
	}

	defer rows.Close()

	// fmt.Println(rows)

	var results []map[string]interface{}
	for rows.Next() {
		var perusahaan string

		err = rows.Scan(&perusahaan)
		if err != nil {
			return nil
		}

		result := make(map[string]interface{})
		result["perusahaan"] = perusahaan

		results = append(results, result)
		// results = result
	}

	if results != nil {
		return c.JSON(200, results)
	}

	return nil

}
