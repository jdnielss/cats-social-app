package usecase

import (
	"fmt"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/repository"
)

type CatUseCase interface {
	Find(q ...string) ([]model.Cat, error)
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

func NewCatUseCase(repo repository.CatRepository) CatUseCase {
	return &catUseCase{repo: repo}
}
