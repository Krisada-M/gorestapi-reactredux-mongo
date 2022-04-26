package controllers

import (
	"server/routes"

	"github.com/gin-gonic/gin"
)

func Expenseroute(rg *gin.RouterGroup) {
	app := rg.Group("/expense")
	app.GET("/", routes.Getexpenses)
	app.GET("/:person", routes.ExpensesByPerson)
	app.POST("/add", routes.AddExpenses)
	app.PATCH("/edit/:id", routes.UpdateExpenses)
	app.PUT("/editall/:id", routes.UpdateAllExpenses)
	app.DELETE("/delete/:id", routes.DeleteExpenses)
}
