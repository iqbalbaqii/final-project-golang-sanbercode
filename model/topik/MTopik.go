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
	coalesce (res.vote, 0) vote,
	coalesce (k.image_src, '') image_src,
	t.created_at,
	t.created_by
		
	FROM topik t 
	INNER join (SELECT * FROM konten where deleted = 0) k on k.id_topik = t.id 
	Left join (select response, count(response) vote from respon group by response) res on res.response = k.id
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
			&polling.Konten.Voting,
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

func GetAll() (result []structs.Topik) {
	var sql = `SELECT * FROM topik`

	raw, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer raw.Close()

	for raw.Next() {
		var polling = structs.Topik{}
		err = raw.Scan(
			&polling.Id,
			&polling.Judul,
			&polling.Pertanyaan,
			&polling.Periode,
			&polling.StartDate,
			&polling.EndDate,
			&polling.CreatedBy,
			&polling.CreatedAt,
			&polling.UpdateAt,
			&polling.Deleted,
			&polling.IsRelease,
			&polling.Target,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, polling)
	}
	return
}

func GetActivePollingBy(id string) (result []structs.Polling) {
	var sql = `SELECT 
	t.id id_topik, 
	t.judul topik,
	t.pertanyaan,
	coalesce (k.id, -1)id_konten,
	coalesce (k.judul, '') judul,
	coalesce (res.vote, 0) vote,
	coalesce (k.image_src, '') image_src,
	t.created_at,
	t.created_by
		
	FROM topik t 
	INNER join (SELECT * FROM konten where deleted = 0) k on k.id_topik = t.id 
	Left join (select response, count(response) vote from respon group by response) res on res.response = k.id
	WHERE t.id = $1
	`

	raw, err := connect.Db.Query(sql, id)
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
			&polling.Konten.Voting,
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
