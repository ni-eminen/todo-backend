package models

type Post struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    body     string    `json:"body"`
    category   string    `json:"category"`
    timestamp uint `json:"timestamp"`
    selected bool `json:"selected"`
    isYesterday bool `json:"yesterday"`
}

type CreateNoteInput struct {
    body     string    `json:"body" binding:"required"`
    category   string    `json:"category" binding:"required"`
    timestamp uint `json:"timestamp" binding:"required"`
    selected bool `json:"selected" binding:"required"`
    isYesterday bool `json:"yesterday" binding:"required"`
}

func CreateNote(c *gin.Context) {
    var input CreateNoteInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    note := models.Post{body: input.body, category: input.category, timestamp: input.timestamp, selected: input.selected, isYesterday: input.isYesterday}
    models.DB.Create(&note)

    c.JSON(http.StatusOK, gin.H{"data": note})
}
