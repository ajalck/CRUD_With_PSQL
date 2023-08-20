//go:build wireinject
// +build wireinject
package di

import (
	"github.com/ajalck/CRUD_With_PSQL/pkg/handlers"
	"github.com/ajalck/CRUD_With_PSQL/pkg/repository"
	"github.com/ajalck/CRUD_With_PSQL/pkg/router"
	"github.com/ajalck/CRUD_With_PSQL/pkg/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeApi(db *gorm.DB) (*router.ServeHTTP, error) {
	wire.Build(
		handlers.NewHandler,
		usecase.NewUseCase,
		repository.NewDB,
		router.NewServeHTTP,
	)
	return &router.ServeHTTP{}, nil
}
