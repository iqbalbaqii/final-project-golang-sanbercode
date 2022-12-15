package listener

import (
	connect "final-project-golang-sanbercode/model"
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
