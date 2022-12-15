package listener

import (
	"final-project-golang-sanbercode/model/listener"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListenerFromTarget(c *gin.Context) {
	var target = c.Param("target")

	var raw = listener.GetListener(target)
	c.JSON(http.StatusOK, gin.H{
		"mesasge": "success",
		"data":    raw,
	})
}
