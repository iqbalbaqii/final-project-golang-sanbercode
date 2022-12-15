package respon

import (
	cons "final-project-golang-sanbercode/model"
	"final-project-golang-sanbercode/structs"
	"fmt"
)

func GetRespon(whereClause string) (res []structs.Respon) {
	sql := fmt.Sprintf("SELECT * FROM respon where %s", whereClause)
	raw, err := cons.Db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer raw.Close()

	for raw.Next() {
		var temp structs.Respon
		err = raw.Scan(&temp.Id, &temp.IdTopik, &temp.Response, &temp.CreatedBy, &temp.CreatedAt, &temp.UpdateAt)
		if err != nil {
			panic(err)
		}
		res = append(res, temp)
	}
	return
}
