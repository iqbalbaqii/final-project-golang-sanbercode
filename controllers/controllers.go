package controllers

import (
	"final-project-golang-sanbercode/structs"
	"strconv"
	"time"
)

var (
	Session *structs.Account
)

func SessionStart(incomming *structs.Account) {
	Session = incomming
}

func GetYear() string {
	year, _, _ := time.Now().Date()
	return strconv.Itoa(year)
}
func GetToday() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FilterMap(data []map[string]string, callback func(map[string]string) bool) []map[string]string {
	var res = []map[string]string{}
	for _, row := range data {
		if callback(row) {
			res = append(res, row)
		}
	}
	return res
}

func Unique(arr []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
