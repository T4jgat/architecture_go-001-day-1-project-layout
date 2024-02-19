package domain

type Group struct {
	ID int `db:"id"`
	// Название группы
	Name string `db:"name" validate:"max=250"`
}
