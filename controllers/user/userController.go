package user

import (
	"encoding/json"
	"final-project-golang-sanbercode/controllers"
	"final-project-golang-sanbercode/model/general"
	"final-project-golang-sanbercode/model/listener"
	"final-project-golang-sanbercode/structs"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddNewUser(c *gin.Context) {
	var validate structs.ValidateUser
	err := c.ShouldBindJSON(&validate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var user structs.User
	var jString, _ = json.Marshal(validate)
	json.Unmarshal([]byte(jString), &user)
	var vGender = validateGender(user.Gender)
	if !vGender {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid gender value",
			"allow": []string{
				"male", "femlae",
			},
		})
		return
	}
	err = general.Insert("\"user\"", map[string]string{
		"username":   user.Username,
		"password":   user.Password,
		"full_name":  user.FullName,
		"gender":     user.Gender,
		"level":      user.Level,
		"email":      user.Email,
		"created_at": controllers.GetToday(),
		"updated_at": controllers.GetToday(),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    user,
		})
		return
	}

	listener.SetListener("ALL", []string{
		user.Username,
	})
	err = listener.SetListener(strings.ToUpper(user.Gender), []string{
		user.Username,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    user,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success add new user",
		"data":    user,
	})

}
func UpdateUser(c *gin.Context) {
	var username = c.Param("username")

	var temp map[string]interface{}

	err := c.ShouldBindJSON(&temp)

	var prepare = map[string]string{}
	for key, row := range temp {
		prepare[key] = row.(string)
	}
	if err != nil {
		return
	}

	prepare["updated_at"] = controllers.GetToday()

	var vGender = validateGender(prepare["gender"])
	if prepare["gender"] != "" && !vGender {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid gender value",
			"allow": []string{
				"male", "femlae",
			},
		})
		return
	}
	err = general.Update("\"user\"", fmt.Sprintf("username = '%s'", username), prepare)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mesasge": fmt.Sprintf("success update. username '%s'", username),
		"change":  prepare,
	})
}

func DeleteUser(c *gin.Context) {
	var id = c.Param("username")
	err := general.Delete("\"user\"", fmt.Sprintf("username = '%s'", id), true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesasge": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mesasge": "success delete username " + id,
	})
}

// ============================
func validateGender(gender string) bool {
	return gender == "female" || gender == "male"
}
