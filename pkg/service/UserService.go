package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xrekson/auction/cmd"
	"github.com/xrekson/auction/pkg/model"
)

func Login(user model.User) (int, gin.H) {
	if len(user.UserName) == 0 || len(user.Password) == 0 {
		return http.StatusBadRequest, gin.H{"error": "Values should not be null"}
	} else {
		userDb, error := cmd.GetUser(user)
		if error != nil {
			return http.StatusNotFound, gin.H{"msg": "User not found!", "error": error}
		}
		if userDb.Password != user.Password {
			return http.StatusBadRequest, gin.H{"error": "Worng password"}
		}
		token, error := cmd.CreateToken(userDb.Name)
		if error == nil {
			return http.StatusOK, gin.H{"msg": "User login success", "token": token}
		} else {
			return http.StatusServiceUnavailable, gin.H{"msg": "Server error!", "error": error}
		}
	}
}
