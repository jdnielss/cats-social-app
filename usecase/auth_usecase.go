package usecase

import (
	"fmt"

	"jdnielss.dev/cats-social-app/model/dto"
	"jdnielss.dev/cats-social-app/repository"
	"jdnielss.dev/cats-social-app/utils/common"
)

type AuthUseCase interface {
	Login(payload dto.LoginRequestDTO) (dto.AuthResponseDTO, error)
}

type authUseCase struct {
	repo repository.AuthRepository
}

func (a *authUseCase) Login(payload dto.LoginRequestDTO) (dto.AuthResponseDTO, error) {
	user, err := a.repo.GetUserData(payload)
	fmt.Println(common.GeneratePasswordHash(payload.Password))
	if err != nil {
		return dto.AuthResponseDTO{}, fmt.Errorf("NOT_FOUND")
	}
	if err := common.CompareHashAndPassword(user.Password, payload.Password); err != nil {
		return dto.AuthResponseDTO{}, fmt.Errorf("INVALID_INPUT")
	}

	var res dto.AuthResponseDTO
	token, err := common.GenerateJWT(user)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}

	res.AccessToken = token
	res.Name = user.Name
	res.Email = user.Email

	return res, nil
}

func NewAuthUseCase(repo repository.AuthRepository) AuthUseCase {
	return &authUseCase{repo: repo}
}
