package web

import (
	"api-contact-form/helper"
	"api-contact-form/model/domain"
)

type ContactResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ContactResponseFromModel(contact *domain.Contact) ContactResponse {
	return ContactResponse{
		ID:        contact.ID,
		Name:      contact.FullName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		Message:   contact.Message,
		CreatedAt: helper.FormatTimeHuman(contact.CreatedAt),
		UpdatedAt: helper.FormatTimeHuman(contact.UpdatedAt),
	}
}
