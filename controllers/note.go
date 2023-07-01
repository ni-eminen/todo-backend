package controllers

import (
	"net/http"

	"github.com/ni-eminen/todo-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateNoteInput struct {
	Id        string `json:"id" binding:"required"`
	Body      string `json:"body" binding:"required"`
	Category  string `json:"category" binding:"required"`
	Timestamp uint   `json:"timestamp" binding:"required"`
	Selected  *bool  `json:"selected" binding:"required"`
}

type UpdateNoteInput struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	Category  string `json:"category"`
	Timestamp uint   `json:"timestamp"`
	Selected  *bool  `json:"selected"`
}

func CreateNote(c *gin.Context) {
	var input CreateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := models.Note{Id: input.Id, Body: input.Body, Category: input.Category, Timestamp: input.Timestamp, Selected: *input.Selected}
	models.DB.Create(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})
}

func FindNotes(c *gin.Context) {
	var notes []models.Note
	models.DB.Find(&notes)

	c.JSON(http.StatusOK, gin.H{"data": notes})
}

func FindNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": note})
}

func UpdateNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.Where("id = ?").First(&note).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var input UpdateNoteInput

	if err := c.ShouldBind(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedNote := models.Note{Id: input.Id, Body: input.Body, Category: input.Category, Timestamp: input.Timestamp, Selected: *input.Selected}

	models.DB.Model(&note).Updates(&updatedNote)
	c.JSON(http.StatusOK, gin.H{"data": note})
}

func DeleteNote(c *gin.Context) {
	var note models.Note
	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&note)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
