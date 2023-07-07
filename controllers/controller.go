package controllers

import (
	"context"
	"csv2mongo/configs"
	"csv2mongo/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var weconnectCollection *mongo.Collection = configs.GetCollection(configs.DB, "weconnect")

func CreateData(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data models.Data

	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error binding data")
	}

	result, err := weconnectCollection.InsertOne(ctx, data)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error saving data to database")
	}

	insertedID := result.InsertedID
	return c.JSON(http.StatusCreated, insertedID)
}

func GetDataByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dataId := c.Param("id")
	var data models.Data

	objId, err := primitive.ObjectIDFromHex(dataId)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	err = weconnectCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&data)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error retrieving data")
	}

	return c.JSON(http.StatusOK, data)
}

func UpdateDataByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dataId := c.Param("id")
	var data models.Data

	objId, err := primitive.ObjectIDFromHex(dataId)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	err = c.Bind(&data)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error binding data")
	}

	_, err = weconnectCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": data})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error updating data")
	}

	return c.String(http.StatusOK, "Successfully updated data")
}

func DeleteDataByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dataId := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(dataId)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	result, err := weconnectCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting data")
	}

	if result.DeletedCount == 0 {
		return c.String(http.StatusNotFound, "Data not found")
	}

	return c.String(http.StatusOK, "Successfully deleted data")
}
