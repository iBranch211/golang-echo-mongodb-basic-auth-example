package security

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whiteListPaths = []string{
	"/favicon.ico",
	"/api",
	"/api/*",
	"/api/v1/signup",
}

func WebSecurityConfig(e *echo.Echo, authValidator *AuthValidator) {
	config := middleware.BasicAuthConfig{
		Validator: authValidator.ValidateCredentials,
		Skipper:   skipAuth,
	}
	e.Use(middleware.BasicAuthWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	// Skip authentication for and signup login requests
	for _, path := range whiteListPaths {
		if path == e.Path() {
			return true
		}
	}
	return false
}
