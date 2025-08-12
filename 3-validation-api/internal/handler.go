package auth

import (
	"api/validation/configs"
	"api/validation/storage"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type AuthHandlerDeps struct {
	*configs.Config
	*storage.Storage
}

type AuthHandler struct {
	*configs.Config
	*storage.Storage
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:  deps.Config,
		Storage: deps.Storage,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Verify())
}

func (handler *AuthHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req EmailRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		// Проверка формата email
		if _, err := mail.ParseAddress(req.Email); err != nil {
			http.Error(w, "invalid email", http.StatusBadRequest)
			return
		}

		// Генерация хеша
		hash := generateHash()
		handler.Storage.Add(req.Email, hash)

		e := email.NewEmail()
		e.From = handler.Auth.Address
		e.To = []string{req.Email}
		e.Subject = "Email verification"
		verifyLink := fmt.Sprintf("http://localhost:8081/verify/%s", hash)
		e.Text = []byte(fmt.Sprintf("Please click the link to verify: %s", verifyLink))
		e.HTML = []byte(fmt.Sprintf("<a href=\"%s\">Click here to verify</a>", verifyLink))

		err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Auth.Email, handler.Auth.Password, "smtp.gmail.com"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "hash %s)", hash)

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

		ok, email := handler.Storage.VerifyAndDelete(hash)
		if ok {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "true (verified for %s)", email)
		} else {
			http.Error(w, "invalid or expired hash", http.StatusUnauthorized)
		}
	}
}

func generateHash() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
