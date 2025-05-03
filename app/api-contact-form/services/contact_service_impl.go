package services

import (
	"api-contact-form/helper"
	"api-contact-form/model/domain"
	"api-contact-form/model/web"
	"api-contact-form/repositories"
	"github.com/go-playground/validator/v10"
)

type ContactServiceImpl struct {
	repository repositories.ContactRepository
	validate   *validator.Validate
}

// instance
func NewContactService(repository repositories.ContactRepository) ContactService {
	return &ContactServiceImpl{
		repository: repository,
		validate:   validator.New(),
	}
}

// createContact
// Memvalidasi input yang diterima dari klien.
// Memetakan data input ke model Contact.
// Menyimpan data ke database menggunakan repository.
func (s *ContactServiceImpl) CreateContact(req *web.ContactRequest) (*domain.Contact, error) {
	// validasi input
	err := s.validate.Struct(req)
	helper.PanicIfError(err)

	// map request to contact form
	contact := &domain.Contact{
		FullName: req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Message:  req.Message,
	}

	//menyimpan ke database
	err = s.repository.Create(contact)
	helper.PanicIfError(err)

	return contact, nil
}

// GetAllContact
// langsung mengembalikan FindAll() dari repository
func (s *ContactServiceImpl) GetAllContact() ([]*domain.Contact, error) {
	return s.repository.FindAll()
}

// GetContactBiId
// langsung mengembalikan FindById() dari repository
func (s *ContactServiceImpl) GetContactById(id uint) (*domain.Contact, error) {
	return s.repository.FindById(id)
}

// UpdateContact
func (s *ContactServiceImpl) UpdateContact(id uint, req *web.ContactRequest) (*domain.Contact, error) {
	// validate input
	err := s.validate.Struct(req)
	helper.PanicIfError(err)

	// mencari id yang lama
	contact, err := s.repository.FindById(id)
	helper.PanicIfError(err)

	// memperbarui data kontak
	contact.FullName = req.Name
	contact.Email = req.Email
	contact.Phone = req.Phone
	contact.Message = req.Message

	// update
	err = s.repository.Update(contact)
	helper.PanicIfError(err)
	return contact, nil
}

// DeleteContact
func (s *ContactServiceImpl) DeleteContact(id uint) error {
	//cari dahulu berdasarkan id
	contact, err := s.repository.FindById(id)
	helper.PanicIfError(err)

	//baru delete
	return s.repository.Delete(contact)
}
