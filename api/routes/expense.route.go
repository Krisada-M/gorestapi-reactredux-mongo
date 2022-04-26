package routes

import (
	"context"
	"fmt"
	"net/http"
	"server/config"
	"server/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	expenseCollection *mongo.Collection = config.MongoCollection(config.DBclient, "expenses")
	expvalidate                         = validator.New()
)

func Getexpenses(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var expenses []bson.M

	result, err := expenseCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = result.All(ctx, &expenses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, expenses)
}

func ExpensesByPerson(c *gin.Context) {
	expperson := c.Param("person")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var expenses []bson.M

	expdata, err := expenseCollection.Find(ctx, bson.M{"person": expperson})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = expdata.All(ctx, &expenses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func AddExpenses(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var expense model.Expensemodel

	if err := c.BindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	expense.ID = primitive.NewObjectID()
	expense.Category = "รายจ่าย"
	expense.Day = config.Day
	expense.Month = config.Month
	expense.Year = config.Year
	expense.Time = config.Time

	validationErr := expvalidate.Struct(expense)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	_, insertErr := expenseCollection.InsertOne(ctx, expense)
	if insertErr != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": expense})
}
