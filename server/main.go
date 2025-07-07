package main

import (
 "fmt"
 "log"
 "net/http"
 "github.com/Despee2k/server/internal/routes"
 "github.com/Despee2k/server/internal/constants"
)

func main() {
	constants.InitEnv()
	r := routes.Setup()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", constants.PORT), r))
}