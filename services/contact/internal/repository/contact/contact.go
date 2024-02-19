package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"context"
)

type ContactRepository interface {
	// CRUD контакта
	CreateContact(ctx context.Context, contact *contact.Contact) error
	GetContact(ctx context.Context, id int) (*contact.Contact, error)
	UpdateContact(ctx context.Context, contact *contact.Contact) error
	DeleteContact(ctx context.Context, id int) error
}
