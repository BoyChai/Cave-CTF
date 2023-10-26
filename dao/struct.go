package dao

import "gorm.io/gorm"

var Dao dao

type dao struct {
}

var db *gorm.DB

// QuestionList 题目表
type QuestionList struct {
	ID       uint   `gorm:"primaryKey;auto_increment" form:"id"`
	Title    string `form:"title"`
	Describe string `form:"describe"`
	Flag     string `form:"flag"`
	Type     string `form:"type"`
	Annex    string `form:"annex"`
	Score    string `form:"score"`
}

// Users 用户表
type Users struct {
	ID    uint   `gorm:"primaryKey;auto_increment" form:"id"`
	Name  string `form:"name"`
	Alias string `form:"alias"`
	Score string `form:"score"`
}

// Ranking 成绩表
type Ranking struct {
	ID            uint   `json:"ID,omitempty;auto_increment" form:"id"`
	User          string `form:"user"`
	Score         string `form:"score"`
	QuestionTitle string `form:"question"`
}
