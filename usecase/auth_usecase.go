package usecase

import (
	"fmt"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/model/dto"
	"jdnielss.dev/cats-social-app/repository"
)

type AuthUseCase interface {
	Login(payload dto.LoginRequestDTO) (model.User, error)
}

type authUseCase struct {
	repo repository.AuthRepository
}

func (a *authUseCase) Login(payload dto.LoginRequestDTO) (model.User, error) {
	user, err := a.repo.GetUserData(payload)
	if err != nil {
		return model.User{}, fmt.Errorf("user with email %s is not registered", payload.Email)
	}
	return user, nil
}

func NewAuthUseCase(repo repository.AuthRepository) AuthUseCase {
	return &authUseCase{repo: repo}
}
