package manager

import "jdnielss.dev/cats-social-app/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	CatRepo() repository.CatRepository
	CourseRepo() repository.CourseRepository
	EnrollmentRepo() repository.EnrollmentRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CourseRepo() repository.CourseRepository {
	return repository.NewCourseRepository(r.infra.Conn())
}

func (r *repoManager) EnrollmentRepo() repository.EnrollmentRepository {
	return repository.NewEnrollmentRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) CatRepo() repository.CatRepository {
	return repository.NewCatRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
