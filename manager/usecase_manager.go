package manager

import "jdnielss.dev/cats-social-app/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CatUseCase() usecase.CatUseCase
	CourseUseCase() usecase.CourseUseCase
	EnrollmentUseCase() usecase.EnrollmentUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CourseUseCase() usecase.CourseUseCase {
	return usecase.NewCourseUseCase(u.repo.CourseRepo())
}

func (u *useCaseManager) EnrollmentUseCase() usecase.EnrollmentUseCase {
	return usecase.NewEnrollmentUseCase(u.repo.EnrollmentRepo(), u.UserUseCase(), u.CourseUseCase())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func (u *useCaseManager) CatUseCase() usecase.CatUseCase {
	return usecase.NewCatUseCase(u.repo.CatRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
