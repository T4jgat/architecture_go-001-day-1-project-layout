package domain

type Contact struct {
	ID int `db:"id"`
	// ФИО, доступное только для чтения
	FullName string `db:"full_name" json:"-"`
	// Фамилия
	FirstName string `db:"first_name"`
	// Имя
	LastName string `db:"last_name"`
	// Отчество
	MiddleName string `db:"middle_name"`
	// Номер телефона
	PhoneNumber string `db:"phone_number"`
}
