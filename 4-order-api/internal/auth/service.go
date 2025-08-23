package auth

import (
	"api/order/internal/user"
	"api/order/pkg/generator"
	"errors"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Connect(phone string) (string, error) {
	var err error
	userEx, _ := service.UserRepository.FindByPhone(phone)
	if userEx == nil {
		userEx = &user.User{
			Phone: phone,
		}
		userEx, err = service.UserRepository.Create(userEx)
		if err != nil {
			return "", err
		}
	}

	sessionId, err := generator.GenerateSessionID(16)
	if err != nil {
		return "", err
	}

	code, err := generator.GenerateCode()
	if err != nil {
		return "", err
	}

	userEx.SessionId = sessionId
	userEx.Code = code
	err = service.UserRepository.Update(userEx)
	if err != nil {
		return "", err
	}

	return userEx.SessionId, nil
}

func (service *AuthService) Confirm(sessionId string, code int) (string, error) {
	userExisted, err := service.UserRepository.FindBySession(sessionId)
	if err != nil {
		return "", errors.New(ErrWrongCredential)
	}

	if userExisted.Code != code {
		return "", errors.New(ErrWrongCredential + " code mismatch")
	}

	return userExisted.SessionId, nil
}
