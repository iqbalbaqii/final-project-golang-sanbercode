package respon

import (
	"final-project-golang-sanbercode/controllers"
	CTopik "final-project-golang-sanbercode/controllers/topik"
	"final-project-golang-sanbercode/model/general"
	MRespon "final-project-golang-sanbercode/model/respon"
	"final-project-golang-sanbercode/model/topik"
	"final-project-golang-sanbercode/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculateRespon(c *gin.Context) {

	var respon structs.Respon
	err := c.ShouldBindJSON(&respon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var whereClause = fmt.Sprintf("id_topik = %s AND created_by = '%s'", respon.IdTopik, controllers.Session.Username)
	var currentRespons = MRespon.GetRespon(whereClause)

	if len(currentRespons) > 0 {
		// update
		err = general.Update("respon", whereClause, map[string]string{
			"response":   respon.Response,
			"updated_at": controllers.GetToday(),
		})
	} else {
		err = general.Insert("respon", map[string]string{
			"id_topik":   respon.IdTopik,
			"response":   respon.Response,
			"created_by": controllers.Session.Username,
			"created_at": controllers.GetToday(),
			"updated_at": controllers.GetToday(),
		})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var getPolling = topik.GetActivePollingBy(respon.IdTopik)
	var CleanPolling = CTopik.CleanPolling(getPolling)

	c.JSON(http.StatusOK, gin.H{
		"message": "your response was saved",
		"data":    CleanPolling,
	})

}
