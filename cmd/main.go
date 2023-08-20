package main

import (
	"sync"

	"github.com/ajalck/CRUD_With_PSQL/pkg/config"
	"github.com/ajalck/CRUD_With_PSQL/pkg/db"
	"github.com/ajalck/CRUD_With_PSQL/pkg/di"
	"github.com/ajalck/CRUD_With_PSQL/pkg/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	wg := &sync.WaitGroup{}
	ch, ch1, ch2 := make(chan interface{}), make(chan interface{}), make(chan interface{})
	wg.Add(3)
	go config.LoadEnv(ch, wg)
	go db.ConnectDB(ch, ch1, ch2, wg)
	go db.ConfigDB(ch1, wg)
	DB := <-ch2
	go di.InitializeApi(DB.(*gorm.DB))
	wg.Wait()
}

func main() {
	port := ":8080"
	r := *gin.Default()
	r.Use(gin.Logger())
	serveHttp := &router.ServeHTTP{Engin: &r}
	serveHttp.Router()
	r.Run(port)
}
