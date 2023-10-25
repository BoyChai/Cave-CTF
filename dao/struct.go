package dao

import "gorm.io/gorm"

var Dao dao

type dao struct {
}

var db *gorm.DB

// QuestionList 题目表
type QuestionList struct {
	ID       uint `gorm:"primaryKey;auto_increment"`
	Title    string
	Describe string
	Flag     string
	Type     string
	Annex    string
	Score    string
}

// Users 用户表
type Users struct {
	ID    uint `gorm:"primaryKey;auto_increment"`
	Name  string
	Score string
}

// Ranking 成绩表
type Ranking struct {
	ID            uint `json:"ID,omitempty;auto_increment"`
	User          string
	Score         string
	QuestionTitle string
}
