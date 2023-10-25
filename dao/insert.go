package dao

// CreateQuestion 创建题目
func (d *dao) CreateQuestion(Question QuestionList) {
	db.Create(&Question)
}

// CreateUsers 创建用户
func (d *dao) CreateUsers(user Users) {
	db.Create(&user)
}

// PutFlag 创建用户
func (d *dao) PutFlag(Flag Ranking) {
	db.Create(&Flag)
}
