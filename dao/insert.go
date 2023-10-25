package dao

// CreateQuestion 创建题目
func (d *dao) CreateQuestion(Question QuestionList) error {
	tx := db.Create(&Question)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// CreateUsers 创建用户
func (d *dao) CreateUsers(user Users) error {
	tx := db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// PutFlag 创建用户
func (d *dao) PutFlag(Flag Ranking) error {
	tx := db.Create(&Flag)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
