package main

import (
	"fmt"
	"go-elastic-engine/app/product/api"
	"go-elastic-engine/app/product/repository"
	"go-elastic-engine/app/product/service"
	"go-elastic-engine/config"
	"log"
	"net/http"
)

func main() {
	loadConfig := config.LoanConfig()
	connDB, err := config.ConnectConfig(loadConfig.Database)

	if err != nil {
		log.Fatalf("error connecting to postgres :%v", err)
		return
	}
	defer connDB.Close()

	fmt.Println("ELASTIC ENGINE PROJECT")
	log.Printf("connected to postgres successfulyy")

	productRepository := repository.NewProductRepository()
	productLogic := service.NewProductLogic(productRepository, connDB)
	productHanlder := api.NewProductHandler(productLogic)
	productRouter := api.NewProductRouter(productHanlder)

	server := http.Server{
		Addr:    "localhost:7654",
		Handler: productRouter,
	}

	fmt.Println("starting server on port 7654...")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
