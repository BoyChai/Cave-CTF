package dao

// DeleteQuestion 创建题目
func (d *dao) DeleteQuestion(id int) {
	question := QuestionList{
		ID: uint(id),
	}
	db.Delete(&question)
}

// DeleteUsers 创建用户
func (d *dao) DeleteUsers(id int) {
	user := Users{
		ID: uint(id),
	}
	db.Delete(&user)

}

// DeleteFlag  创建用户
func (d *dao) DeleteFlag(id int) {
	flag := Ranking{
		ID: uint(id),
	}
	db.Delete(&flag)

}
