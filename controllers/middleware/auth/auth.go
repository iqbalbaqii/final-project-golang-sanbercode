package middleware

import (
	cons "final-project-golang-sanbercode/controllers"
	CUser "final-project-golang-sanbercode/controllers/user"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	found, data := CUser.CheckLogin(user, password)
	if hasAuth && found {
		cons.SessionStart(data)
	} else {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.JSON(403, gin.H{
			"message": "unauthorized",
		})
		return
	}
}
