package dao

import "gorm.io/gorm"

var db *gorm.DB

type QuestionList struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Describe string
	Flag     string
	Type     string
	Annex    string
	Score    string
}

type Users struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Ranking struct {
	ID    uint `json:"ID,omitempty"`
	User  string
	Score string
}
