package repositories

import (
	"api-contact-form/model/domain"
)

// kontrak contact_repository_impl
type ContactRepository interface {
	Create(contact *domain.Contact) error
	FindAll() ([]*domain.Contact, error)
	FindById(id uint) (*domain.Contact, error)
	Update(contact *domain.Contact) error
	Delete(contact *domain.Contact) error
}
