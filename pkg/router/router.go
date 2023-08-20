package router

import (
	"github.com/ajalck/CRUD_With_PSQL/pkg/handlers"
	"github.com/gin-gonic/gin"
)

type ServeHTTP struct {
	Handler *handlers.Handlers
	Engin   *gin.Engine
}

func NewServeHTTP(handler *handlers.Handlers) *ServeHTTP {
	return &ServeHTTP{Handler: handler}
}

func (s *ServeHTTP) Router() {
	s.Engin.POST("/insert", s.Handler.InsertStudentData)
}
