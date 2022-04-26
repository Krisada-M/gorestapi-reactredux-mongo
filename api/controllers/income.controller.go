package controllers

import (
	"server/routes"

	"github.com/gin-gonic/gin"
)

func Incomeroute(rg *gin.RouterGroup) {
	app := rg.Group("/income")
	app.GET("/", routes.GetIncome)
	app.GET("/:person", routes.IncomeByPerson)
	app.POST("/add", routes.AddIncome)
	app.PATCH("/edit/:id", routes.UpdateIncome)
	app.PUT("/editall/:id", routes.UpdateAllIncome)
	app.DELETE("/delete/:id", routes.DeleteIncome)
}
