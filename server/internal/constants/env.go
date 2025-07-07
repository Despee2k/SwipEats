package constants

import (
	"github.com/Despee2k/server/internal/config"
	"strconv"
)

var PORT int

func InitEnv() {
	config.LoadEnv()
	portStr := config.GetEnv("PORT", "8080")

	port, err := strconv.Atoi(portStr)

	if err != nil {
		panic("Invalid PORT value in .env file: " + portStr)
	}

	PORT = port
}