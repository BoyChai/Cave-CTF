package dao

import "fmt"

// FindQuestionList 查询题目
func (d *dao) FindQuestionList() []QuestionList {
	var list []QuestionList
	r := db.Find(&list)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
	}
	return list
}

// FindUser 查询所有用户
func (d *dao) FindUser() []Users {
	var list []Users
	r := db.Find(&list)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
	}
	return list
}

// FindRanking 查询成绩
func (d *dao) FindRanking() []Ranking {
	var list []Ranking
	r := db.Find(&list)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
	}
	return list
}
