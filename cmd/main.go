package main

import (
	"fmt"
	"github.com/beuus39/product/internal/app"
	"github.com/beuus39/product/internal/config"
	productGrpcPresenter "github.com/beuus39/product/internal/ports/grpc/handler"
	productGrpc "github.com/beuus39/product/internal/ports/grpc/server"
	queue2 "github.com/beuus39/product/internal/ports/queue"
	"github.com/beuus39/product/internal/repository/product"
	"github.com/beuus39/product/pkg/nats"
	"os"
)

func main() {
	natsConf := nats.NatsConfig{
		Url: "nats://localhost:4222",
	}
	productQueue := queue2.NewProductQueue(natsConf)

	fmt.Println("Welcome to codebase")
	conf := config.LoadEnv()
	repo, _ := product.NewMongoRepository(conf.URL, conf.DB, conf.Timeout)
	service := app.NewProductService(repo)
	/*handler := rest.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products/{code}", handler.Get)
	r.Post("/products", handler.Post)
	r.Delete("/products/{code}", handler.Delete)
	r.Get("/products", handler.GetAll)
	r.Put("/products", handler.Put)*/

	// Start GRPC
	productGrpcHandler := productGrpcPresenter.NewGrpcProductHandler(service, productQueue)
	grpcServer, err := productGrpc.NewGrpcServer(productGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}

	err = grpcServer.Serve(conf.GrpcPort)
	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}
	//log.Fatal(http.ListenAndServe(conf.ServerPort, r))
}