// CORS (Cross-Origin Resource Sharing) adalah mekanisme keamanan yang digunakan oleh browser untuk mengontrol
// bagaimana sumber daya (resources) di server dapat diakses oleh aplikasi web dari domain lain.
package middleware

import (
	"api-contact-form/helper"
	"github.com/gin-contrib/cors"
)

func SetupCors() cors.Config {
	return cors.Config{
		// CORS Config
		AllowOrigins:     helper.ParseEnvList("CORS_ALLOWED_ORIGINS"),
		AllowMethods:     helper.ParseEnvList("CORS_ALLOWED_METHODS"),
		AllowHeaders:     helper.ParseEnvList("CORS_ALLOWED_HEADERS"),
		AllowCredentials: helper.GetEnvBool("CORS_ALLOW_CREDENTIALS", false),
		MaxAge:           12 * 60 * 60, // 12 hours
	}
}
