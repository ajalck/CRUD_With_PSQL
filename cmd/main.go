package main

import (
	"sync"

	"github.com/ajalck/CRUD_With_PSQL/pkg/config"
	"github.com/ajalck/CRUD_With_PSQL/pkg/db"
)

func init() {
	wg := &sync.WaitGroup{}
	ch, ch1 := make(chan interface{}), make(chan interface{})
	wg.Add(3)
	go config.LoadEnv(ch, wg)
	go db.ConnectDB(ch, ch1, wg)
	go db.ConfigDB(ch1, wg)
	wg.Wait()
}

func main() {

}
