package auth

import (
	"api/order/configs"
	"api/order/pkg/jwt"
	"api/order/pkg/req"
	"api/order/pkg/resp"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/connect", handler.Connect())
	router.HandleFunc("POST /auth/confirm", handler.Confirm())
}

func (handler *AuthHandler) Connect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[ConnectRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sessionId, err := handler.AuthService.Connect(payload.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		data := ConnectResponse{
			SessionId: sessionId,
		}

		resp.Json(w, data, http.StatusAccepted)
	}
}

func (handler *AuthHandler) Confirm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[VerifyRequest](w, r)
		if err != nil {
			return
		}

		sessionId, err := handler.AuthService.Confirm(payload.SessionId, payload.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(sessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := VerifyResponse{
			Token: token,
		}

		resp.Json(w, data, http.StatusOK)
	}
}
