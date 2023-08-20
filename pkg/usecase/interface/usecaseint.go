package interfaces

type UseCase interface {
	InsertStudentData()
	UpdateStudentData()
	DeleteStudentData()
	ReadStudentData()
	FilterStudentData()
}
