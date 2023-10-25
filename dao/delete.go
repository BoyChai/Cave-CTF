package dao

// DeleteQuestion 创建题目
func (d *dao) DeleteQuestion(id int) error {
	question := QuestionList{
		ID: uint(id),
	}
	tx := db.Delete(&question)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteUsers 创建用户
func (d *dao) DeleteUsers(id int) error {
	user := Users{
		ID: uint(id),
	}
	tx := db.Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteFlag  创建用户
func (d *dao) DeleteFlag(id int) error {
	flag := Ranking{
		ID: uint(id),
	}
	tx := db.Delete(&flag)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
