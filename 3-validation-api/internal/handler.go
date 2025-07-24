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
		e.From = handler.Auth.Address
		e.To = []string{"test@example.com"}
		e.Bcc = []string{"test_bcc@example.com"}
		e.Cc = []string{"test_cc@example.com"}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Auth.Email, handler.Auth.Password, "smtp.gmail.com"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (handler *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлечение hash из URL
		hash := r.PathValue("hash")
		if hash == "" {
			http.Error(w, "missing hash", http.StatusBadRequest)
			return
		}

		// Пример логики верификации (здесь просто проверка длины или конкретного значения)
		if isValidHash(hash) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Hash verified successfully")
		} else {
			http.Error(w, "invalid or expired hash", http.StatusUnauthorized)
		}
	}
}

func isValidHash(hash string) bool {
	return len(hash) == 32
}
