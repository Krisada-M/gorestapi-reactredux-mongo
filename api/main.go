package main

import (
	"log"
	"net/http"
	"os"
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func welcome(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><title>api data</title><h1>api runing</h1></html>"))
}

// func init() {
// 	gin.SetMode(gin.ReleaseMode)
// }

func main() {
	config.Envload()
	config.DBconnect()

	//port
	port := os.Getenv("PORT")
	if port == "" {
		port = "6504"
	}

	app := gin.New()

	// Gin middleware
	app.Use(gin.Logger())
	app.Use(cors.Default())

	app.GET("/", welcome)
	base := app.Group("/")
	controllers.Incomeroute(base)
	controllers.Expenseroute(base)

	log.Fatal(app.Run(":" + port))

}
