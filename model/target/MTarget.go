package target

import (
	connect "final-project-golang-sanbercode/model"
	"final-project-golang-sanbercode/structs"
)

func GetAllTarget() (res []structs.Target) {
	sql := "SELECT * FROM target"

	raw, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer raw.Close()

	for raw.Next() {
		var temp structs.Target
		err = raw.Scan(&temp.Id, &temp.Target, &temp.Description, &temp.CreatedAt, &temp.UpdateAt)
		if err != nil {
			panic(err)
		}
		res = append(res, temp)
	}
	return
}

func GetTargetByName(target string) []structs.Target {
	sql := "SELECT * FROM target WHERE target = $1"

	raw, err := connect.Db.Query(sql, target)
	if err != nil {
		panic(err)
	}
	defer raw.Close()

	res := []structs.Target{}
	for raw.Next() {
		var temp structs.Target
		err = raw.Scan(&temp.Id, &temp.Target, &temp.Description, &temp.CreatedAt, &temp.UpdateAt)
		if err != nil {
			panic(err)
		}
		res = append(res, temp)
	}

	return res
}
