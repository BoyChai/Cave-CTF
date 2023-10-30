package dao

import "gorm.io/gorm"

var Dao dao

type dao struct {
}

var db *gorm.DB

// QuestionList 题目表
type QuestionList struct {
	ID       uint   `gorm:"primaryKey;auto_increment" form:"id"`
	Title    string `form:"title" binding:"required"`
	Describe string `form:"describe" binding:"required"`
	Flag     string `form:"flag" binding:"required"`
	Type     string `form:"type" binding:"required"`
	Annex    string `form:"annex" binding:"required"`
	Score    string `form:"score" binding:"required"`
}

// Users 用户表
type Users struct {
	ID uint `gorm:"primaryKey;auto_increment" form:"id"`
	//Name  string `form:"name" binding:"required" gorm:"unique"`
	Name  string `form:"name" binding:"required"`
	Alias string `form:"alias" binding:"required" gorm:"unique"`
}

// // Ranking 成绩表
//
//	type Ranking struct {
//		ID            uint   `json:"ID,omitempty;auto_increment" form:"id"`
//		User          string `form:"user" binding:"required"`
//		Score         string `form:"score" binding:"required"`
//		QuestionTitle string `form:"question" binding:"required"`
//	}
//
// Ranking 成绩表
type Ranking struct {
	ID            uint   `json:"ID,omitempty;auto_increment" form:"id"`
	User          string `form:"user" binding:"required"`
	Score         string `form:"score" binding:"required"`
	QuestionTitle string `form:"question" binding:"required" gorm:"unique_index"`
}
