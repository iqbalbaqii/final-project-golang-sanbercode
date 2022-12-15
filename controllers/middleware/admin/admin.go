package admin

import (
	"final-project-golang-sanbercode/controllers"

	"github.com/gin-gonic/gin"
)

func HasTobeSuper(c *gin.Context) {
	var current = *controllers.Session
	if current.Level == "0" {

	} else {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.JSON(403, gin.H{
			"message": "unauthorized",
		})
		return
	}
}
