package dao

import (
	"fmt"
)

// FindQuestion 查询单个题目
func (d *dao) FindQuestion(id int) (error, QuestionList) {
	var Question QuestionList
	tx := db.First(&Question, id)
	if tx.Error != nil {
		return tx.Error, QuestionList{}
	}
	return nil, Question
}

// FindQuestionList 查询题目
func (d *dao) FindQuestionList() (error, []QuestionList) {
	var list []QuestionList
	r := db.Find(&list)
	if r.Error != nil {
		return r.Error, nil
	}
	return nil, list
}

// FindUsers 查询所有用户
func (d *dao) FindUsers() (error, []Users) {
	var list []Users
	r := db.Find(&list)
	if r.Error != nil {
		return r.Error, nil
	}
	return nil, list
}

// FindUser 查询单个用户
func (d *dao) FindUser(name string, alias string) (error, Users) {
	var list Users
	r := db.Where("name = ? AND alias = ?", name, alias).First(&list)
	if r.Error != nil {
		return r.Error, Users{}
	}
	return nil, list
}

// FindRankings 查询成绩
func (d *dao) FindRankings() (error, []Ranking) {
	var list []Ranking
	r := db.Find(&list)
	if r.Error != nil {
		return r.Error, nil
	}
	return nil, list
}

// FindRanking 查询单个成绩
func (d *dao) FindRanking(name string, title string) (error, Ranking) {

	var ranking Ranking
	tx := db.Where("user = ? AND question_title=?", name, title).First(&ranking)
	if tx.Error != nil {
		fmt.Println(1)
		return tx.Error, Ranking{}
	}

	return nil, ranking
}

// FindPersonalRanking 查询个人获取的成绩列表
func (d *dao) FindPersonalRanking(name string) (error, []Ranking) {
	var ranking []Ranking
	tx := db.Where("user = ?", name).Find(&ranking)
	if tx.Error != nil {
		return tx.Error, []Ranking{}
		fmt.Println(tx.Error)
	}
	return nil, ranking
}
