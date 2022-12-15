package konten

import (
	"encoding/json"
	"final-project-golang-sanbercode/model/general"
	"final-project-golang-sanbercode/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeContent(c *gin.Context) {

	var validationKonten structs.ValidationKonten
	err := c.ShouldBindJSON(&validationKonten)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var jsonString, _ = json.Marshal(validationKonten)
	var byteData = []byte(jsonString)
	var konten structs.Konten
	json.Unmarshal(byteData, &konten)

	err = general.Insert("konten", map[string]string{
		"id_topik":   konten.IdTopik,
		"judul":      konten.Judul,
		"keterangan": konten.Keterangan,
		"image_src":  konten.ImageSrc,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": konten,
	})
}

func DeleteKonten(c *gin.Context) {
	var id = c.Param("id")
	err := general.Delete("konten", "id = "+id, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mesasge": "success delete konten id = " + id,
	})
}

func UpdateKonten(c *gin.Context) {
	var id = c.Param("id")

	var temp map[string]interface{}

	err := c.ShouldBindJSON(&temp)

	var prepare = map[string]string{}
	for key, row := range temp {
		prepare[key] = row.(string)
	}
	if err != nil {
		return
	}
	err = general.Update("konten", "id = "+id, prepare)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mesasge": "success update",
		"data":    temp,
	})
}
