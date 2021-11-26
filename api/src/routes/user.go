package routes

import (
	"api/src/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type create_user_parameters struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Create_user(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody create_user_parameters

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}

		if len(requestBody.Email) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		if len(requestBody.Name) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		if len(requestBody.Password) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		// db.First(&product, "code = ?", "D42") // find product with code D42
		db.Create(&types.User{Email: requestBody.Email, Name: requestBody.Name, Password: requestBody.Password})
		c.JSON(http.StatusOK, gin.H{"user": requestBody.Name})
	}
}
