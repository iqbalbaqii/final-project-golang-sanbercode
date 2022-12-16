package main

import (
	CKonten "final-project-golang-sanbercode/controllers/konten"
	CList "final-project-golang-sanbercode/controllers/listener"
	admin "final-project-golang-sanbercode/controllers/middleware/admin"
	auth "final-project-golang-sanbercode/controllers/middleware/auth"
	CRespon "final-project-golang-sanbercode/controllers/respon"
	CTarget "final-project-golang-sanbercode/controllers/target"
	CTopik "final-project-golang-sanbercode/controllers/topik"
	CUser "final-project-golang-sanbercode/controllers/user"
	connect "final-project-golang-sanbercode/model"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	connect.Start()
	// defer connect.Db.Close()

	var router = gin.Default()
	var root = router.Group("/polling", auth.BasicAuth)

	root.GET("/aktif", CTopik.GetActivePolling)

	root.GET("/topik", CTopik.GetTopik)
	root.GET("/topik/:id_topik", CTopik.GetTopikID)
	root.POST("/topik", CTopik.AddTopik)
	root.PUT("/topik/:id", CTopik.UpdateTopik)
	root.DELETE("/topik/:id", CTopik.DeleteTopik)

	root.POST("/konten", CKonten.MakeContent)
	root.PUT("/konten/:id", CKonten.UpdateKonten)
	root.DELETE("/konten/:id", CKonten.DeleteKonten)

	root.GET("/user", CUser.GetUser)
	root.POST("/user", admin.HasTobeSuper, CUser.AddNewUser)
	root.PUT("/user/:username", admin.HasTobeSuper, CUser.UpdateUser)
	root.DELETE("/user/:username", admin.HasTobeSuper, CUser.DeleteUser)

	root.POST("/respon", CRespon.CalculateRespon)

	root.POST("/target", admin.HasTobeSuper, CTarget.Insert)
	root.GET("/target", admin.HasTobeSuper, CTarget.Get)

	root.GET("/listener/:target", admin.HasTobeSuper, CList.GetListenerFromTarget)

	router.Run(":" + os.Getenv("PORT"))

}
