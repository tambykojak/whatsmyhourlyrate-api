package main

import (
	"os"
	"strconv"

	"github.com/tambykojak/whatsmyhourlyrate-api/server"
)

func main() {
	server := server.Server{Port: getPort()}
	server.Initialize()
}

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return 3000
	}

	return port
}
