package manager

import "jdnielss.dev/cats-social-app/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	CatUseCase() usecase.CatUseCase
}

type useCaseManager struct {
	repo RepoManager
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
