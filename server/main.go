package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/routes"
	"github.com/SwipEats/SwipEats/server/internal/utils"
)

func main() {
	constants.InitEnv()

	db.ConnectDatabase()
	db.MigrateModels()

	r := routes.Setup(utils.CreateGroupSessionService())

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", constants.PORT), r))
}