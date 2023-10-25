package dao

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

// FindUser 查询所有用户
func (d *dao) FindUser() (error, []Users) {
	var list []Users
	r := db.Find(&list)
	if r.Error != nil {
		return r.Error, nil
	}
	return nil, list
}

// FindRanking 查询成绩
func (d *dao) FindRanking() (error, []Ranking) {
	var list []Ranking
	r := db.Find(&list)
	if r.Error != nil {
		return r.Error, nil
	}
	return nil, list
}
