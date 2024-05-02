package manager

import "jdnielss.dev/cats-social-app/usecase"

type UseCaseManager interface {
	AuthUseCase() usecase.AuthUseCase
	CatUseCase() usecase.CatUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.repo.AuthRepo())
}

func (u *useCaseManager) CatUseCase() usecase.CatUseCase {
	return usecase.NewCatUseCase(u.repo.CatRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
