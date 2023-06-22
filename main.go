package main

import (
    "github.com/ni-eminen/todo-backend/models"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()  // new!

    router.POST("/notes", controllers.CreateNote)  // here!

    router.Run("localhost:8080")
}