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
	incomeCollection *mongo.Collection = config.MongoCollection(config.DBclient, "incomes")
	incvalidate                        = validator.New()
)

func GetIncome(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var income []bson.M
	result, err := incomeCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	if err = result.All(ctx, &income); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, income)
}

func IncomeByPerson(c *gin.Context) {
	incperson := c.Param("person")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var income []bson.M
	incdata, err := incomeCollection.Find(ctx, bson.M{"person": incperson})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = incdata.All(ctx, &income); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, income)
}

func AddIncome(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var income model.IncExpmodel

	if err := c.BindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	income.ID = primitive.NewObjectID()
	income.Category = "รายรับ"
	income.Date.Day = config.Day
	income.Date.Month = config.Month
	income.Date.Year = config.Year
	income.Date.Time = config.Time

	validationErr := incvalidate.Struct(income)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	_, insertErr := incomeCollection.InsertOne(ctx, income)
	if insertErr != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": income})
}

func UpdateIncome(c *gin.Context) {
	incomeid := c.Param("id")
	mid, _ := primitive.ObjectIDFromHex(incomeid)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var income model.IncExpmodel

	if err := c.BindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := incomeCollection.UpdateOne(ctx, bson.M{"_id": mid},
		bson.D{bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "productname", Value: income.Productname},
			bson.E{Key: "purchase", Value: income.Purchase},
		}}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, result.ModifiedCount)
}

func UpdateAllIncome(c *gin.Context) {
	incomeid := c.Param("id")
	mid, _ := primitive.ObjectIDFromHex(incomeid)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var income model.IncExpmodel

	if err := c.BindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := incvalidate.Struct(income)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, err := incomeCollection.ReplaceOne(ctx, bson.M{"_id": mid},
		bson.M{
			"category":    "รายรับ",
			"person":      income.Person,
			"productname": income.Productname,
			"purchase":    income.Purchase,
			"date": bson.M{
				"day":   config.Day,
				"month": config.Month,
				"year":  config.Year,
				"time":  config.Time,
			},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, result.ModifiedCount)
}

func DeleteIncome(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	incomeid := c.Param("id")
	mid, _ := primitive.ObjectIDFromHex(incomeid)
	result, err := incomeCollection.DeleteOne(ctx, bson.M{"_id": mid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, result.DeletedCount)
}
