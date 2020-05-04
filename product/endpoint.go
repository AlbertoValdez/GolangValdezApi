package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//GetProductByIDRequest estructura
type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

func makeGetProductByIDEndPoint(s Service) endpoint.Endpoint {
	getProductByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(getProductByIDRequest)
		product, err := s.GetProductByID(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIDEndPoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {

			panic(err)
		}

		return result, nil
	}

	return getProductsEndPoint
}

func makeAddProductEndPoint(s Service) endpoint.Endpoint {

	addProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productID, error := s.InsertProducts(&req)
		if error != nil {
			panic(error)
		}

		return productID, nil
	}

	return addProductEndPoint

}
