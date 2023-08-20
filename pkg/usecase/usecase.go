package usecase

import (
	repoInt "github.com/ajalck/CRUD_With_PSQL/pkg/repository/interfaces"
	usecaseInt "github.com/ajalck/CRUD_With_PSQL/pkg/usecase/interface"
)

type UseCase struct {
	Repo repoInt.Repo
}

func NewUseCase(repo repoInt.Repo) usecaseInt.UseCase {
	return &UseCase{repo}
}

func (u *UseCase) InsertStudentData() {

}

func (u *UseCase) UpdateStudentData() {

}

func (u *UseCase) DeleteStudentData() {

}

func (u *UseCase) ReadStudentData() {

}

func (u *UseCase) FilterStudentData() {

}
