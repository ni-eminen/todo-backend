package models

type Note struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Body      string `json:"body"`
	Category  string `json:"category"`
	Timestamp uint   `json:"timestamp"`
	Selected  bool   `json:"selected"`
}
