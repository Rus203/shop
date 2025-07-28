package main

import (
	"github.com/Rus203/shop/cmd/api"
)

func main() {
	apiServer := api.NewApiServer()

	apiServer.Run()
}