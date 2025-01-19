package main

import (
	"log"
	"project/api-gateway/infra"
	"project/api-gateway/routes"
)

func main() {
	var err error

	var ctx *infra.ServiceContext
	if ctx, err = infra.NewServiceContext(); err != nil {
		log.Fatal("can't init service context %w", err)
	}

	routes.NewRoutes(*ctx)
}
