package topik

import (
	cons "final-project-golang-sanbercode/controllers"
	"final-project-golang-sanbercode/model/general"
	Mtarget "final-project-golang-sanbercode/model/target"
	MTopik "final-project-golang-sanbercode/model/topik"
	"final-project-golang-sanbercode/structs"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetActivePolling(c *gin.Context) {
	rows := MTopik.GetActivePolling()
	if len(rows) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no active polling found",
		})
		return
	}
	var res = CleanPolling(rows)

	c.JSON(http.StatusOK, gin.H{
		"message": "success load active polling",
		"data":    res,
	})
}

func GetTopik(c *gin.Context) {
	var today = cons.GetToday()
	rows := MTopik.GetAll()
	if len(rows) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no topik found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success load topik polling",
		"data":    rows,
	})
	_ = today
}

func AddTopik(c *gin.Context) {
	var topik structs.Topik
	err := c.ShouldBindJSON(&topik)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	var vTarget = validateTarget(topik.Target)
	if !vTarget {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unlisted value target",
			"allow":   Mtarget.GetAllTarget(),
		})
		return
	}

	var vStartDate = validateDate(topik.StartDate)
	var vEndDate = validateDate(topik.EndDate)
	fmt.Println(vStartDate, vEndDate)
	if !vEndDate || !vStartDate {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "uncorrect star_date or end_date format",
			"allow":   "YYYY-MM-DD",
		})
		return
	}

	err = general.Insert("topik", map[string]string{
		"judul":      topik.Judul,
		"pertanyaan": topik.Pertanyaan,
		"start_date": topik.StartDate,
		"end_date":   topik.EndDate,
		"created_at": cons.GetToday(),
		"updated_at": cons.GetToday(),
		"created_by": cons.Session.Username,
		"periode":    cons.GetYear(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": topik,
	})
}

func UpdateTopik(c *gin.Context) {
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
	err = general.Update("topik", "id = "+id, prepare)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mesasge": "success update topik " + id,
		"data":    temp,
	})
}

func DeleteTopik(c *gin.Context) {
	var id = c.Param("id")

	var temp map[string]interface{}

	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": "something wrong",
		})
		return
	}
	var prepare = map[string]bool{}
	for key, row := range temp {
		prepare[key] = row.(bool)
	}

	err = general.Delete("topik", "id = "+id, prepare["hard_delete"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mesasge": "success delete",
	})
}

func CleanPolling(rows []structs.Polling) (res []gin.H) {
	var toMap []map[string]string
	var topiks []string
	for _, row := range rows {
		idTopik := strconv.Itoa(int(row.Topik.Id))
		idKonten := strconv.Itoa(int(row.Konten.Id))
		topiks = append(topiks, idTopik)
		var temp = map[string]string{
			"id_topik":    idTopik,
			"judul":       row.Topik.Judul,
			"pertanyaan":  row.Topik.Pertanyaan,
			"konten":      row.Konten.Judul,
			"id_konten":   idKonten,
			"jumlah_vote": row.Konten.Voting,
			"image_url":   row.Konten.ImageSrc,
			"created_at":  row.Topik.CreatedAt,
			"created_by":  row.Topik.CreatedBy,
		}
		toMap = append(toMap, temp)
	}
	topiks = cons.Unique(topiks)

	for _, topik := range topiks {
		var temp = cons.FilterMap(toMap, func(row map[string]string) bool {
			return row["id_topik"] == topik
		})
		var konten = []gin.H{}
		for _, row := range temp {
			if row["konten"] == "" {
				continue
			}
			konten = append(konten, gin.H{
				"id":          row["id_konten"],
				"image_url":   row["image_url"],
				"judul":       row["konten"],
				"jumlah_vote": row["jumlah_vote"],
			})
		}

		var current = temp[0]
		var prepareData = gin.H{
			"judul":      current["judul"],
			"id":         current["id_topik"],
			"pertanyaan": current["pertanyaan"],
			"created_at": current["created_at"],
			"created_by": current["created_by"],
			"konten":     konten,
		}
		res = append(res, prepareData)
	}
	return
}

func GetTopikID(c *gin.Context) {

	rows := MTopik.GetActivePollingBy(c.Param("id_topik"))
	if len(rows) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "polling not found",
		})
		return
	}
	var res = CleanPolling(rows)

	c.JSON(http.StatusOK, gin.H{
		"message": "success load active polling",
		"topik":   res,
	})
}

// ====================================================
func validateTarget(data string) bool {
	var target = Mtarget.GetTargetByName(data)
	fmt.Println(target)
	return len(target) != 0
}

func validateDate(date string) bool {
	matched, _ := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, date)

	return matched
}
