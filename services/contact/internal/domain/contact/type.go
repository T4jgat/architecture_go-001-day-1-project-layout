package contact

import "fmt"

type Contact struct {
	ID int `db:"id"`
	// ФИО, доступное только для чтения
	FullName string `db:"full_name" json:"-"`
	// Фамилия
	FirstName string `db:"first_name"`

	Email string `db:"email"`

	Phone string `db:"phone"` // Имя
}

// GetFullName - метод для получения ФИО
func (c *Contact) GetFullName() string {
	return fmt.Sprintf("%s %s %s", c.FirstName, c.Email, c.Phone)
}
