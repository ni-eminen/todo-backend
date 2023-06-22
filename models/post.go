package models

import "time"

type Post struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    body     string    `json:"body"`
    category   string    `json:"category"`
    timestamp uint `json:"timestamp"`
    selected bool `json:"selected"`
    isYesterday bool `json:"yesterday"`
}
