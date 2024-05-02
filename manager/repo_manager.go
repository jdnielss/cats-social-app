package manager

import "jdnielss.dev/cats-social-app/repository"

type RepoManager interface {
	AuthRepo() repository.AuthRepository
	CatRepo() repository.CatRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) AuthRepo() repository.AuthRepository {
	return repository.NewAuthRepository(r.infra.Conn())
}

func (r *repoManager) CatRepo() repository.CatRepository {
	return repository.NewCatRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
