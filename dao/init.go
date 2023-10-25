package dao

import "fmt"

func AutoTables() {
	var err error
	// 题目表
	err = db.AutoMigrate(&QuestionList{})
	if err != nil {
		fmt.Println(err)
	}
	// 成员表
	err = db.AutoMigrate(&Users{})
	if err != nil {
		fmt.Println(err)
	}
	// 分数表
	err = db.AutoMigrate(&Ranking{})
	if err != nil {
		fmt.Println(err)
	}
}
