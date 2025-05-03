// MainHandler memberikan informasi dasar tentang API kepada pengguna atau developer.

package handlers

import (
	"api-contact-form/model/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainHandler struct {
}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

func (h *MainHandler) MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "SUCCESS",
		Message: "API Contact is Runnning",
	})
}
