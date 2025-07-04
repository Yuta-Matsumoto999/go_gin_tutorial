package main

import (
	"github.com/Yuta-Matsumoto999/go_gin_tutorial/controllers"
	"github.com/Yuta-Matsumoto999/go_gin_tutorial/models"
	"github.com/Yuta-Matsumoto999/go_gin_tutorial/repositories"
	"github.com/Yuta-Matsumoto999/go_gin_tutorial/services"
	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "Item 1", Price: 100, Description: "Description 1", SoldOut: false},
		{ID: 2, Name: "Item 2", Price: 200, Description: "Description 2", SoldOut: true},
		{ID: 3, Name: "Item 3", Price: 300, Description: "Description 3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	router := gin.Default()
	router.GET("/items", itemController.FindAll)
	router.GET("/items/:id", itemController.FindById)
	router.POST("/items", itemController.Create)
	router.PUT("/items/:id", itemController.Update)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
