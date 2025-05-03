package repositories

import (
	"api-contact-form/model/domain"
	"gorm.io/gorm"
	"time"
)

// ContactRepositoryImpl adalah implementasi dari interface ContactRepository.
type ContactRepositoryImpl struct {
	db *gorm.DB
}

// NewContactRepository membuat instance baru dari ContactRepository.
func NewContactRepository(db *gorm.DB) ContactRepository {
	return &ContactRepositoryImpl{db: db}
}

// create
func (repository *ContactRepositoryImpl) Create(contact *domain.Contact) error {
	return repository.db.Create(contact).Error
}

// FindAll
func (repository *ContactRepositoryImpl) FindAll() ([]*domain.Contact, error) {
	var contacts []*domain.Contact
	err := repository.db.Where("deleted_at = ?", "0000-00-00 00:00:00").Find(&contacts).Error
	return contacts, err
}

// FindById
func (repository *ContactRepositoryImpl) FindById(id uint) (*domain.Contact, error) {
	var contact domain.Contact
	err := repository.db.Where("id = ? AND deleted_at = ?", id, "0000-00-00 00:00:00").First(&contact).Error
	return &contact, err
}

// Update
func (repository *ContactRepositoryImpl) Update(contact *domain.Contact) error {
	return repository.db.Save(contact).Error
}

// Soft Delete (tidak menghapus db)
func (repository *ContactRepositoryImpl) Delete(contact *domain.Contact) error {
	contact.DeletedAt = time.Now()
	return repository.db.Save(contact).Error
}
