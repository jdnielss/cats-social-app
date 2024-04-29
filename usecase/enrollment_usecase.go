package usecase

import (
	"fmt"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/model/dto"
	"jdnielss.dev/cats-social-app/repository"
)

type EnrollmentUseCase interface {
	RegisterNewEnrollment(payload dto.EnrollmentRequestDto) (model.Enrollment, error)
	FindById(id string) (model.Enrollment, error)
}

type enrollmentUseCase struct {
	repo     repository.EnrollmentRepository
	userUC   UserUseCase
	courseUC CourseUseCase
}

func (e *enrollmentUseCase) FindById(id string) (model.Enrollment, error) {
	return e.repo.Get(id)
}

func (e *enrollmentUseCase) RegisterNewEnrollment(payload dto.EnrollmentRequestDto) (model.Enrollment, error) {
	var newEnrollmentDetail []model.EnrollmentDetail
	course, err := e.courseUC.FindById(payload.CourseId)
	if err != nil {
		return model.Enrollment{}, err
	}

	for _, v := range payload.Users {
		user, err := e.userUC.FindById(v)
		if err != nil {
			return model.Enrollment{}, err
		}
		newEnrollmentDetail = append(newEnrollmentDetail, model.EnrollmentDetail{User: user})
	}

	newEnrollment := model.Enrollment{
		Course:            course,
		EnrollmentDetails: newEnrollmentDetail,
	}

	enrollment, err := e.repo.Create(newEnrollment)
	if err != nil {
		return model.Enrollment{}, fmt.Errorf("failed to create enrollment: %s", err.Error())
	}

	return enrollment, nil
}

func NewEnrollmentUseCase(
	repo repository.EnrollmentRepository,
	userUC UserUseCase,
	courseUC CourseUseCase,
) EnrollmentUseCase {
	return &enrollmentUseCase{
		repo:     repo,
		userUC:   userUC,
		courseUC: courseUC,
	}
}
