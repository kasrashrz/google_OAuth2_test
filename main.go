package main

import (
	"github.com/kasrashrz/google_OAuth2_test/internal/configs"
	"github.com/kasrashrz/google_OAuth2_test/internal/services"
	"net/http"
)

func main() {
	configs.InitializeViper()

	// Initialize Logger across the application
	//logger.InitializeZapCustomLogger()

	// Initialize Oauth2 Services
	services.InitializeOauthGoogle()

	// Routes for the application
	http.HandleFunc("/", services.HandleMain)
	http.HandleFunc("/login-gl", services.HandleGoogleLogin)
	http.HandleFunc("/callback-gl", services.CallBackFromGoogle)

	//logger.("Started running on http://localhost:" + "9090")
	http.ListenAndServe(":9090", nil)
}
