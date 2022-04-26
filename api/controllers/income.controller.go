package controllers

import (
	"server/routes"

	"github.com/gin-gonic/gin"
)

func Incomeroute(rg *gin.RouterGroup) {
	app := rg.Group("/income")
	app.GET("/allincome", routes.GetIncome)
	app.GET("/:person", routes.IncomeByPerson)
	app.POST("/addincome", routes.AddIncome)

}
