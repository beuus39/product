package main

import (
	"fmt"
	"github.com/beuus39/product/internal/app"
	"github.com/beuus39/product/internal/config"
	productGrpcPresenter "github.com/beuus39/product/internal/ports/grpc/handler"
	productGrpc "github.com/beuus39/product/internal/ports/grpc/server"
	"github.com/beuus39/product/internal/ports/rest"
	"github.com/beuus39/product/internal/repository/product"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
)

const grpcDefaultPort = 3002
func main() {
	fmt.Println("Welcome to codebase")
	conf, _ := config.NewConfig("C:/Users/ADMIN/go/src/beu.us/codebase/internal/config/config.yaml")
	repo, _ := product.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	service := app.NewProductService(repo)
	handler := rest.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products/{code}", handler.Get)
	r.Post("/products", handler.Post)
	r.Delete("/products/{code}", handler.Delete)
	r.Get("/products", handler.GetAll)
	r.Put("/products", handler.Put)

	// Start GRPC
	productGrpcHandler := productGrpcPresenter.NewGrpcProductHandler(service)
	grpcServer, err := productGrpc.NewGrpcServer(productGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}

	err = grpcServer.Serve(grpcDefaultPort)
	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}
	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}