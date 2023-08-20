package repository

import (
	repoInt "github.com/ajalck/CRUD_With_PSQL/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type DataBase struct {
	DB *gorm.DB
}

func NewDB(DB *gorm.DB) repoInt.Repo {
	return &DataBase{DB}
}

func (db *DataBase) InsertStudentData() {

}
func (db *DataBase) UpdateStudentData() {

}
func (db *DataBase) DeleteStudentData() {

}
func (db *DataBase) ReadStudentData() {

}
func (db *DataBase) FilterStudentData() {

}
