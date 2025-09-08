package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware provides a pre-configured CORS setup for Gin.
// It allows requests from specified origins, methods, and headers.
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4321",     // Your API's own origin
			"http://127.0.0.1:4321",     // Another common localhost variant
			"https://hoppscotch.io",     // Hoppscotch's main domain
			"https://app.hoppscotch.io", // Another common Hoppscotch domain
			// Add your frontend origins here in production: e.g., "https://your-frontend.com"
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400, // Cache preflight requests for 24 hours
	})
}
