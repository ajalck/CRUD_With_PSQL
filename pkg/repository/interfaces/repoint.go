package interfaces

type Repo interface {
	InsertStudentData()
	UpdateStudentData()
	DeleteStudentData()
	ReadStudentData()
	FilterStudentData()
}
