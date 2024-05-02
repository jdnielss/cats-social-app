package usecase

import (
	"fmt"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/model/dto"
	"jdnielss.dev/cats-social-app/repository"
	"jdnielss.dev/cats-social-app/utils/common"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
	Register(payload model.User) (dto.AuthResponseDTO, error)
	GetByEmail(email string) (bool, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	// menambahkan validasi custom dipersilahkan
	// Misalnya kita buat pesan yang lebih informatif
	// ID tidak ditemukan
	user, err := u.repo.Get(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) GetByEmail(email string) (bool, error) {
	_, err := u.repo.GetByEmail(email)

	if err != nil {
		return true, fmt.Errorf("user with email %s already exists", email)
	}

	return false, nil
}

func (u *userUseCase) Register(payload model.User) (dto.AuthResponseDTO, error) {

	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}
	payload.Password = newPassword
	return u.repo.Create(payload)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
