package usecase

import (
	"architecture_go/services/contact/internal/domain"
	"context"
)

type ContactUseCase interface {
	// CreateContact contacts
	CreateContact(ctx context.Context, contact *domain.Contact) error
	GetContact(ctx context.Context, id int) (*domain.Contact, error)
	UpdateContact(ctx context.Context, contact *domain.Contact) error
	DeleteContact(ctx context.Context, id int) error

	// CreateGroup groups
	CreateGroup(ctx context.Context, group *domain.Group) error
	GetGroup(ctx context.Context, id int) (*domain.Group, error)
	AddContactToGroup(ctx context.Context, contactID, groupID int) error
}
