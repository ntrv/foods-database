package models

type Food struct {
	Model
	Name string `json:name`
}

func (f Food) AllFoods() ([]Food, error) {
	db, err := NewDBConnect()
	if err != nil {
		return nil, err
	}

	var foods []Food
	db.Find(&foods)
	return foods, nil
}
