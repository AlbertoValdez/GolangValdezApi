package product

import (
	"context"

	"github.com/GolangValdezApi/helper"
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

type updateProductRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID string
}

type getBestSellersRequest struct {
}

func makeGetProductByIDEndPoint(s Service) endpoint.Endpoint {
	getProductByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(getProductByIDRequest)
		product, err := s.GetProductByID(&req)
		helper.Catch(err)
		return product, nil
	}
	return getProductByIDEndPoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		helper.Catch(err)

		return result, nil
	}

	return getProductsEndPoint
}

func makeAddProductEndPoint(s Service) endpoint.Endpoint {

	addProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productID, error := s.InsertProducts(&req)
		helper.Catch(error)

		return productID, nil
	}

	return addProductEndPoint

}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {

	updateProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		r, err := s.UpdateProducts(&req)

		helper.Catch(err)
		return r, nil
	}

	return updateProductEndPoint
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {

	delateProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(deleteProductRequest)
		r, err := s.DeleteProducts(&req)

		helper.Catch(err)
		return r, nil
	}
	return delateProductEndPoint
}

func makeBestSellersEndPoint(s Service) endpoint.Endpoint {

	getBestSellersEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {

		result, err := s.GetBestSellers()
		helper.Catch(err)
		return result, err
	}

	return getBestSellersEndPoint

}
