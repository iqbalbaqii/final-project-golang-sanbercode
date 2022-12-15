package general

import (
	"errors"
	connect "final-project-golang-sanbercode/model"
	"fmt"
)

func Insert(table string, data map[string]string) error {
	if len(data) == 0 {
		return errors.New("data pass cannot be empty")
	}
	var field = ""
	var value = ""
	for key, val := range data {
		field += fmt.Sprintf("%s, ", key)
		value += fmt.Sprintf("'%s', ", val)
	}
	field = field[:len(field)-2]
	value = value[:len(value)-2]
	sql := fmt.Sprintf("INSERT Into %s (%s) VALUES (%s)", table, field, value)
	errs := connect.Db.QueryRow(sql)
	return errs.Err()
}

func Update(table string, where string, data map[string]string) error {
	fmt.Println(table, where, data)
	var sentence = ""
	for key, val := range data {
		sentence += fmt.Sprintf("%s = '%s', ", key, val)
	}
	sentence = sentence[:len(sentence)-2]
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, sentence, where)
	errs := connect.Db.QueryRow(sql)
	return errs.Err()
}

func Delete(table string, where string, hardDelete bool) error {
	var sql string
	if hardDelete {
		sql = fmt.Sprintf("DELETE  FROM %s WHERE %s", table, where)
	} else {
		return Update(table, where, map[string]string{
			"deleted": "1",
		})
	}
	errs := connect.Db.QueryRow(sql)
	return errs.Err()
}
