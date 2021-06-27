package server

import (
	"fmt"
	productGrpcPresenter "github.com/beuus39/product/internal/ports/grpc/handler"
	pb "github.com/beuus39/product/pkg/grpc/product"
	"google.golang.org/grpc"
	"net"
)
type Server struct {
	productHandler *productGrpcPresenter.GrpcProductHandler
}

func (s *Server) Serve(port uint) error {
	address := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	pb.RegisterProductServiceServer(server, s.productHandler)
	err = server.Serve(l)

	if err != nil {
		return err
	}
	fmt.Printf("Product GRPC Server running on PORT %d", port)
	return nil
}

func NewGrpcServer(productGrpcHandler *productGrpcPresenter.GrpcProductHandler) (*Server, error) {
	return &Server{
		productHandler: productGrpcHandler,
	}, nil
}