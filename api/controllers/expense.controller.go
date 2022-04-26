package controllers

import (
	"server/routes"

	"github.com/gin-gonic/gin"
)

func Expenseroute(rg *gin.RouterGroup) {
	app := rg.Group("/expense")
	app.GET("/allexpense", routes.Getexpenses)
	app.GET("/:person", routes.ExpensesByPerson)
	app.POST("/addexpense", routes.AddExpenses)
}
