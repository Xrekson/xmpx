package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xrekson/auction/cmd"
	"github.com/xrekson/auction/pkg/model"
	"github.com/xrekson/auction/pkg/service"
)

func Loginhandeler(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, res := service.Login(user)
	c.JSON(status, res)
}

func Alluserhandeler(c *gin.Context) {
	users, errors := cmd.GetAllusers()
	if len(errors) == 0 {
		c.JSON(http.StatusOK, gin.H{"usersData": users})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Users not found!", "errors": errors})
	}
}
