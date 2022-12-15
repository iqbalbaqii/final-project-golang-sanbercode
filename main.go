package main

import (
	CKonten "final-project-golang-sanbercode/controllers/konten"
	CTopik "final-project-golang-sanbercode/controllers/topik"
	CUser "final-project-golang-sanbercode/controllers/user"
	connect "final-project-golang-sanbercode/model"

	"github.com/gin-gonic/gin"
)

func main() {
	connect.Start()
	// defer connect.Db.Close()

	var router = gin.Default()
	var root = router.Group("/polling")
	root.GET("/aktif", CTopik.GetActivePolling)

	root.POST("/topik", CTopik.AddTopik)
	root.PUT("/topik/:id", CTopik.UpdateTopik)
	root.DELETE("/topik/:id", CTopik.DeleteTopik)

	root.POST("/konten", CKonten.MakeContent)
	root.PUT("/konten/:id", CKonten.UpdateKonten)
	root.DELETE("/konten/:id", CKonten.DeleteKonten)

	root.POST("/user", CUser.AddNewUser)
	root.PUT("/user/:username", CUser.UpdateUser)
	root.DELETE("/user/:username", CUser.DeleteUser)

	router.Run("localhost:8080")
}
