package auth

import (
	"api/validation/configs"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Verify())
}

func (handler *AuthHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{"test@example.com"}
		e.Bcc = []string{"test_bcc@example.com"}
		e.Cc = []string{"test_cc@example.com"}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (handler *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify from Handler!")
	}
}
