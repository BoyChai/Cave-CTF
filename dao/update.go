package dao

func (d *dao) Score(id int, Score string) error {
	user := Users{
		ID: uint(id),
		//Score: Score,
	}
	tx := db.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
