package main

import (
	CTopik "final-project-golang-sanbercode/controllers/topik"
	connect "final-project-golang-sanbercode/model"
)

func main() {
	connect.Start()
	defer connect.Db.Close()

	CTopik.GetActivePolling()
}
