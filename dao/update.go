package dao

func (d *dao) Score(id int, Score string) {
	user := Users{
		ID:    uint(id),
		Score: Score,
	}
	db.Save(&user)
}
