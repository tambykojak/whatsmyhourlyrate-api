package main

import (
	"os"
	"strconv"
)

func main() {
	app := app{port: getPort()}
	app.initialize()
}

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return 3000
	}

	return port
}
