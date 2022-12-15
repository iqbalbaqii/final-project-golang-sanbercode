package topik

import (
	cons "final-project-golang-sanbercode/controllers"
	MTopik "final-project-golang-sanbercode/model/topik"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetActivePolling(c *gin.Context) {
	var today = cons.GetToday()
	rows := MTopik.GetActivePolling()

	var toMap []map[string]string
	var topiks []string
	for _, row := range rows {
		idTopik := strconv.Itoa(int(row.Topik.Id))
		topiks = append(topiks, idTopik)
		var temp = map[string]string{
			"id_topik":   idTopik,
			"judul":      row.Topik.Judul,
			"pertanyaan": row.Topik.Pertanyaan,
			"konten":     row.Konten.Judul,
		}
		toMap = append(toMap, temp)
	}
	topiks = cons.Unique(topiks)
	var res [][]map[string]string
	for _, topik := range topiks {
		var temp = cons.FilterMap(toMap, func(row map[string]string) bool {
			return row["id_topik"] == topik
		})
		res = append(res, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success load active polling",
		"data":    res,
	})
	_ = today
}
