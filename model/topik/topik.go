package topik

import (
	connect "final-project-golang-sanbercode/model"
	"final-project-golang-sanbercode/structs"
)

func GetActivePolling() (result []structs.Polling) {
	var sql = `SELECT t.id, t.judul topik, t.pertanyaan, coalesce (k.judul, '') judul  FROM topik t 
	LEFT join konten k on k.id_topik = t.id `

	rows, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var polling = structs.Polling{}
		err = rows.Scan(&polling.Topik.Id, &polling.Topik.Judul, &polling.Topik.Pertanyaan, &polling.Konten.Judul)

		if err != nil {
			panic(err)
		}
		result = append(result, polling)
	}
	return
}
