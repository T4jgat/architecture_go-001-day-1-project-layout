package http

import (
	"architecture_go/services/contact/internal/domain/contact"
	"context"
	"net/http"
)

type ContactHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type ContactRepository interface {
	Create(ctx context.Context, contact contact.Contact) error
	Get(ctx context.Context, id int) (contact.Contact, error)
	Update(ctx context.Context, contact contact.Contact) error
	Delete(ctx context.Context, id int) error
}
