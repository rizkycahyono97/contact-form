package handlers

import (
	"api-contact-form/model/web"
	"api-contact-form/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ContactHandler struct {
	service services.ContactService
}

// instance
func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{service: service}
}

// CreatedContact Handler
func (h *ContactHandler) CreateContact(c *gin.Context) {
	var req web.ContactRequest

	//Bind the JSON payload to the ContactRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Use the service layer to create a new contact.
	contact, err := h.service.CreateContact(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Respond with the created contact and a success message.
	c.JSON(http.StatusCreated, web.ApiResponse{
		Code:    "CREATED",
		Message: "Contact Created Successfully",
		Data:    web.ContactResponseFromModel(contact),
	})
}

// GetAllContact Handler
func (h *ContactHandler) GetContacts(c *gin.Context) {
	// Fetch all contacts using the service layer.
	contacts, err := h.service.GetAllContact()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	//  Konversi Model ke Format Respons
	var contactResponses []web.ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, web.ContactResponseFromModel(contact))
	}

	// Respond with the list of contacts.
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Contacts Retrivied Successfully",
		Data:    contactResponses,
	})
}

// GetContactById
func (h *ContactHandler) GetContact(c *gin.Context) {
	// ambil id dari parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid ID format. ID must be a number.",
			Data:    nil,
		})
		return
	}

	// Menggunakan service layer untuk mencari kontak berdasarkan ID.
	contact, err := h.service.GetContactById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, web.ApiResponse{
			Code:    "NOT_FOUND",
			Message: "Contact not found",
			Data:    nil,
		})
		return
	}

	// Mengembalikan data kontak dalam format JSON jika ditemukan
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Contact Retrivied Successfully",
		Data:    web.ContactResponseFromModel(contact),
	})
}

// UpdateContct
func (h *ContactHandler) UpdateContact(c *gin.Context) {
	// ambil id dari parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "INVALID_ID",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	//Bind the JSON payload to the ContactRequest struct
	var req web.ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Use the service layer to update the contact.
	contact, err := h.service.UpdateContact(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Success
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Contact Updated Successfully",
		Data:    web.ContactResponseFromModel(contact),
	})
}

// Delete
func (h *ContactHandler) DeleteContact(c *gin.Context) {
	// ambil id dari parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Use the service layer to update the contact.
	err = h.service.DeleteContact(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Berhasil
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "Contact Deleted Successfully",
		Data:    nil,
	})
}
