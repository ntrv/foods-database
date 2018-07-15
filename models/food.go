package models

type Food struct {
	Model
	Name string `json:name form:name validate:required`
}

func (f Food) AllFoods() ([]Food, error) {
	db, err := NewDBConnect()
	if err != nil {
		return nil, err
	}

	var foods []Food
	if err := db.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func (f Food) RegisterFood() error{
	db, err := NewDBConnect()
	if err != nil {
		return err
	}
	return db.Debug().Create(&f).Error
}
