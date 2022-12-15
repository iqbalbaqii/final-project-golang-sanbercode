package main

import (
	CTopik "final-project-golang-sanbercode/controllers/topik"
	connect "final-project-golang-sanbercode/model"

	"github.com/gin-gonic/gin"
)

func main() {
	connect.Start()
	defer connect.Db.Close()

	var router = gin.Default()
	var root = router.Group("/polling")
	root.GET("/aktif", CTopik.GetActivePolling)

	router.Run("localhost:8080")
}
