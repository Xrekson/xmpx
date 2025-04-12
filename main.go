package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xrekson/auction/cmd"
	"github.com/xrekson/auction/pkg/middleware"
	"github.com/xrekson/auction/web"
)

func main() {
	os.Setenv("JWTKEY", "ebde68cba15731310e0cf345d7468cc99561d02696eb9cf8016759e7ac68a2fe")
	os.Setenv("DATABASE_URL", "postgres://app:Thundera@190@localhost:5432/app")
	cmd.CreateSchema()
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello, World!",
		})
	})
	router.POST("/login", web.Loginhandeler)
	protected := router.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/all/user", web.Alluserhandeler)
		protected.GET("/all/listing", web.Alllistinghandeler)
	}

	router.Run(":19090")
}
