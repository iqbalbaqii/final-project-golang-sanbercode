package listener

import (
	connect "final-project-golang-sanbercode/model"
	"final-project-golang-sanbercode/structs"
	"fmt"
	"strings"
)

func SetListener(target string, username []string) error {
	var statement []string
	for _, user := range username {
		var combo = fmt.Sprintf("('%s', '%s')", user, target)
		statement = append(statement, combo)
	}

	var clause = strings.Join(statement, " ")
	fmt.Println(clause)
	sql :=
		`
	insert into listener (username, target)
	values 
	` + clause

	errs := connect.Db.QueryRow(sql)
	return errs.Err()
}

func GetListener(target string) (res []structs.Listener) {
	sql := fmt.Sprintf("SELECT * FROM listener where target = '%s'", target)
	raw, err := connect.Db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer raw.Close()

	for raw.Next() {
		var temp structs.Listener
		err = raw.Scan(&temp.Id, &temp.Target, &temp.Username)
		if err != nil {
			panic(err)
		}
		res = append(res, temp)
	}
	return
}
