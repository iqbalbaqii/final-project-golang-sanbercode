package topik

import (
	connect "final-project-golang-sanbercode/model"
	"final-project-golang-sanbercode/structs"
)

func GetActivePolling() (result []structs.Polling) {
	var sql = `SELECT 
		t.id id_topik, 
		t.judul topik,
		t.pertanyaan,
		coalesce (k.id, -1)id_konten,
		coalesce (k.judul, '') judul,
		coalesce (k.image_src, '') image_src,
		t.created_at,
		t.created_by
		  
	FROM topik t 
	LEFT join (SELECT * FROM konten where deleted = 0) k on k.id_topik = t.id 
	WHERE t.deleted = 0 AND t.is_release = 1
	`

	raw, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer raw.Close()

	for raw.Next() {
		var polling = structs.Polling{}
		err = raw.Scan(
			&polling.Topik.Id,
			&polling.Topik.Judul,
			&polling.Topik.Pertanyaan,
			&polling.Konten.Id,
			&polling.Konten.Judul,
			&polling.Konten.ImageSrc,
			&polling.Topik.CreatedAt,
			&polling.Topik.CreatedBy,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, polling)
	}
	return
}
