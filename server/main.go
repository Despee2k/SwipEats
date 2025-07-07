package main

import (
 "fmt"
 "log"
 "net/http"
 "github.com/Despee2k/server/internal/routes"
 "github.com/Despee2k/server/internal/constants"
 "github.com/Despee2k/server/internal/db"
)

func main() {
	constants.InitEnv()

	db.ConnectDatabase()
	db.MigrateModels()

	r := routes.Setup()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", constants.PORT), r))
}