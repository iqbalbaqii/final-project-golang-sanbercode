package target

import (
	"final-project-golang-sanbercode/controllers"
	"final-project-golang-sanbercode/model/general"
	"final-project-golang-sanbercode/model/target"
	"final-project-golang-sanbercode/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Insert(c *gin.Context) {
	var target structs.Target
	err := c.ShouldBindJSON(&target)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	general.Insert("target", map[string]string{
		"target":      target.Target,
		"description": target.Description,
		"created_at":  controllers.GetToday(),
		"updated_at":  controllers.GetToday(),
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    target,
	})
}

func Get(c *gin.Context) {
	var raw = target.GetAllTarget()
	c.JSON(http.StatusOK, gin.H{
		"mesasge": "success",
		"data":    raw,
	})
}
