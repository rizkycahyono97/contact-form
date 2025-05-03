package services

import (
	"api-contact-form/model/domain"
	"api-contact-form/model/web"
)

// Kontrak ContactServiceImpl
type ContactService interface {
	CreateContact(req *web.ContactRequest) (*domain.Contact, error)
	GetAllContact() ([]*domain.Contact, error)
	GetContactById(id uint) (*domain.Contact, error)
	UpdateContact(id uint, req *web.ContactRequest) (*domain.Contact, error)
	DeleteContact(id uint) error
}
