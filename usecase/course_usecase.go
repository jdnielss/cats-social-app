package usecase

import (
	"fmt"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/repository"
)

type CourseUseCase interface {
	FindById(id string) (model.Course, error)
}

type courseUseCase struct {
	repo repository.CourseRepository
}

func (c *courseUseCase) FindById(id string) (model.Course, error) {
	course, err := c.repo.Get(id)
	if err != nil {
		return model.Course{}, fmt.Errorf("course with ID %s not found", id)
	}
	return course, nil
}

func NewCourseUseCase(repo repository.CourseRepository) CourseUseCase {
	return &courseUseCase{repo: repo}
}
