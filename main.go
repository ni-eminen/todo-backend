package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ni-eminen/todo-backend/controllers"
	"github.com/ni-eminen/todo-backend/models"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	models.ConnectDatabase()

	router.POST("/notes", controllers.CreateNote) // here!
	router.GET("/notes", controllers.FindNotes)
	router.GET("/notes/:id", controllers.FindNote)
	router.PATCH("/ntoes/:id", controllers.UpdateNote)
	router.DELETE("/posts/:id", controllers.DeleteNote)

	router.Run("localhost:8080")
}
