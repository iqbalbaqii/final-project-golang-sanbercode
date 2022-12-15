package konten

import (
	cons "final-project-golang-sanbercode/model"
	"final-project-golang-sanbercode/structs"
	"fmt"
)

func GetKonten(whereClause string) (res []structs.Konten) {
	sql := fmt.Sprintf("SELECT * FROM konten where %s AND deleted = 0", whereClause)
	raw, err := cons.Db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer raw.Close()

	for raw.Next() {
		var temp structs.Konten
		err = raw.Scan(&temp.Id, &temp.IdTopik, &temp.Judul, &temp.Keterangan, &temp.ImageSrc, &temp.Deleted)
		if err != nil {
			panic(err)
		}
		res = append(res, temp)
	}
	return
}
