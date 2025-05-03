//HealthHandler digunakan untuk memastikan bahwa aplikasi berjalan dengan baik,
//terutama untuk kebutuhan monitoring dan otomatisasi.

package handlers

import (
	"api-contact-form/model/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, web.ApiResponse{
		Code: "Success",
		Data: "API is working",
	})
}
