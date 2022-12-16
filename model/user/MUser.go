package user

import (
	connect "final-project-golang-sanbercode/model"
	"final-project-golang-sanbercode/structs"
)

func GetUserBy(username string) []structs.Account {
	sql := "SELECT username, password, email, level FROM \"user\" WHERE username = $1"

	raw, err := connect.Db.Query(sql, username)
	if err != nil {
		panic(err)
	}
	defer raw.Close()

	res := []structs.Account{}
	for raw.Next() {
		var temp structs.Account
		err = raw.Scan(&temp.Username, &temp.Password, &temp.Email, &temp.Level)
		if err != nil {
			panic(err)
		}
		res = append(res, temp)
	}

	return res
}

func GetAll() []structs.User {
	sql := "SELECT username, password, full_name, email, level FROM \"user\""

	raw, err := connect.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer raw.Close()

	res := []structs.User{}
	for raw.Next() {
		var temp structs.User
		err = raw.Scan(&temp.Username, &temp.Password, &temp.FullName, &temp.Email, &temp.Level)
		if err != nil {
			panic(err)
		}
		res = append(res, temp)
	}

	return res
}
