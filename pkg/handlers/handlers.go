package handlers

import (
	usecaseInt "github.com/ajalck/CRUD_With_PSQL/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	usecase usecaseInt.UseCase
}

func NewHandler(usecase usecaseInt.UseCase) *Handlers {
	return &Handlers{usecase}
}

func (h *Handlers) InsertStudentData(ctx *gin.Context) {

}

func (h *Handlers) UpdateStudentData(ctx *gin.Context) {

}

func (h *Handlers) DeleteStudentData(ctx *gin.Context) {

}

func (h *Handlers) ReadStudentData(ctx *gin.Context) {

}

func (h *Handlers) FilterStudentData(ctx *gin.Context) {

}
