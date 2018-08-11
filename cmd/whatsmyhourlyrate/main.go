package main

import (
	"os"
	"strconv"

	"github.com/tambykojak/whatsmyhourlyrate-api/internal/whatsmyhourlyrate"
)

func main() {
	app := whatsmyhourlyrate.App{Port: getPort()}
	app.Initialize()
}

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return 3000
	}

	return port
}
