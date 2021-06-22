package handler

import (
	"context"
	"github.com/beuus39/product/internal/app"
	"github.com/beuus39/product/internal/domain"
	pb "github.com/beuus39/product/pkg/grpc/product"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

type GrpcProductHandler struct {
	service app.Service
}

func (h *GrpcProductHandler) FindByCategory(request *pb.ProductQueryRequest, server pb.ProductService_FindByCategoryServer) error {
	log.Fatal("implement me")
	return nil
}

func NewGrpcProductHandler(service app.Service) *GrpcProductHandler {
	return &GrpcProductHandler{service: service}
}

func (h *GrpcProductHandler) FindByID(ctx context.Context, arg *pb.ProductQueryRequest) (*pb.ProductResponse, error)  {
	id := arg.ID

	productResult := <-h.service.FindProduct(int(id))

	if productResult.Error != nil {
		return nil, status.Error(codes.InvalidArgument, productResult.Error.Error())
	}

	product, ok := productResult.Result.(*domain.Product)
	if !ok {
		err := errors.New("Result is not product")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	stock, _ := strconv.Atoi(product.Stock.String())
	price, _ := strconv.ParseFloat(product.Price.String(), 32)

	productResponse := &pb.ProductResponse{
		ID: int32(product.ID),
		CategoryID: int32(product.CategoryID),
		Name: product.Name,
		Description: product.Description,
		Image: product.Image,
		Stock: int32(stock),
		Price: price,
	}
	return productResponse, nil
}

func (h *GrpcProductHandler) FindAll(arg *pb.ProductQueryRequest, stream pb.ProductService_FindAllServer) error  {
	productResult := <-h.service.FindAllProducts()
	if productResult.Error != nil {
		return status.Error(codes.InvalidArgument, productResult.Error.Error())
	}
	products, ok := productResult.Result.([]*domain.Product)

	if !ok {
		err := errors.New("Result is Not Products")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	for _, product := range products {
		stock, _ := strconv.Atoi(product.Stock.String())
		price, _ := strconv.ParseFloat(product.Price.String(), 32)

		productResponse := &pb.ProductResponse{
			ID: int32(product.ID),
			CategoryID: int32(product.CategoryID),
			Name: product.Name,
			Description: product.Description,
			Image: product.Image,
			Stock: int32(stock),
			Price: price,
		}

		if err := stream.Send(productResponse); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}
