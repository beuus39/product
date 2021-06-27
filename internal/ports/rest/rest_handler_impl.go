package rest

import (
	"encoding/json"
	"github.com/beuus39/product/internal/app"
	"github.com/beuus39/product/internal/domain"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type handler struct {
	productService app.Service
}

func (h handler) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(request, "id"))

	productResult := <-h.productService.FindProduct(id)

	if productResult.Error != nil {
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	product, ok := productResult.Result.(*domain.Product)

	if !ok {
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(writer).Encode(product)
}

func (h handler) Post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	p := &domain.Product{}

	err := json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	savedProduct := <-h.productService.SaveProduct(p)

	if savedProduct.Error != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(savedProduct.Result)
}

func (h handler) Put(writer http.ResponseWriter, request *http.Request) {
	p := &domain.Product{}

	err := json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	updatedProduct := <-h.productService.UpdateProduct(p)

	if updatedProduct.Error != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(updatedProduct.Result)
}

func (h handler) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(request, "id"))
	deletedProduct := <-h.productService.DeleteProduct(id)

	if deletedProduct.Error != nil {
		log.Fatal(deletedProduct.Error)
	}
}

func (h handler) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	productResult := <-h.productService.FindAllProducts()

	if productResult.Error != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	products, ok := productResult.Result.([]*domain.Product)
	if !ok {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(products)
}

func NewHandler(productService app.Service) ProductHandler {
	return &handler{productService: productService}
}

