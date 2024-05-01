package usecase

import (
	"fmt"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/model/dto"
	"jdnielss.dev/cats-social-app/repository"
)

type CatUseCase interface {
	Find(q ...string) ([]model.Cat, error)
	Create(payload dto.CatRequestDTO) (dto.CreateCatResponseDTO, error)
}

type catUseCase struct {
	repo repository.CatRepository
}

func (c *catUseCase) Find(q ...string) ([]model.Cat, error) {
	cat, err := c.repo.Get(q)

	if err != nil {
		fmt.Printf(`Error Repo %s`, err)
		return []model.Cat{}, fmt.Errorf("Error")
	}

	return cat, nil
}

func (c *catUseCase) Create(payload dto.CatRequestDTO) (dto.CreateCatResponseDTO, error) {
	cat, err := c.repo.Create(payload)

	return cat, err
}

func NewCatUseCase(repo repository.CatRepository) CatUseCase {
	return &catUseCase{repo: repo}
}
