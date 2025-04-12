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
	if errors == nil {
		c.JSON(http.StatusOK, gin.H{"usersData": users})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Users not found!", "errors": errors})
	}
}

func CreateListing(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"list": "User has created Listing!"})
}

func Alllistinghandeler(c *gin.Context) {
	listings, errors := cmd.GetAllListings()
	if errors == nil {
		c.JSON(http.StatusOK, gin.H{"listingData": listings})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Listings not found!", "errors": errors})
	}
}

