package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"context"
)

type ContactUseCase interface {
	// CRUD контакта
	CreateContact(ctx context.Context, contact *contact.Contact) error
	GetContact(ctx context.Context, id int) (*contact.Contact, error)
	UpdateContact(ctx context.Context, contact *contact.Contact) error
	DeleteContact(ctx context.Context, id int) error
}

func (u *contactUseCase) CreateContact(ctx context.Context, contact contact.Contact) error {
	// Бизнес-логика перед созданием
	// ...

	// Сохранение контакта
	err := u.repo.Create(ctx, contact)
	if err != nil {
		return err
	}

	// Бизнес-логика после создания
	// ...

	return nil
}

func (u *contactUseCase) GetContact(ctx context.Context, id int) (contact.Contact, error) {
	// Получение контакта из репозитория
	contact, err := u.repo.Get(ctx, id)
	if err != nil {
		return contact.Contact{}, err
	}

	// Бизнес-логика после получения
	// ...

	return contact, nil
}

func (u *contactUseCase) UpdateContact(ctx context.Context, contact contact.Contact) error {
	// Бизнес-логика перед обновлением
	// ...

	// Обновление контакта
	err := u.repo.Update(ctx, contact)
	if err != nil {
		return err
	}

	// Бизнес-логика после обновления
	// ...

	return nil
}
