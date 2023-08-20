package main

import (
	"github.com/ajalck/CRUD_With_PSQL/pkg/config"
	"github.com/ajalck/CRUD_With_PSQL/pkg/db"
)

func init(){
	c := config.LoadEnv()
	db.ConnectDB(c)
}

func main() {
	
}
