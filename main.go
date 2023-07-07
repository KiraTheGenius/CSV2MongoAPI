package main

import (
	"csv2mongo/configs"
	"csv2mongo/controllers"
	"csv2mongo/data"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDB()

	// Run code below to add data to database once!!!
	res, err := data.AddData()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	e := echo.New()
	e.POST("/data/", controllers.CreateData)
	e.GET("/data/:id", controllers.GetDataByID)
	e.PUT("/data/:id", controllers.UpdateDataByID)
	e.DELETE("/data/:id", controllers.DeleteDataByID)
	e.Logger.Fatal(e.Start(":8080"))
}
