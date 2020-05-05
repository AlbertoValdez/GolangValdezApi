package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

//MakeHTTPHandler funcion
func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()
	GetProductByIDHandler := kithttp.NewServer(makeGetProductByIDEndPoint(s), getProductByIDRequesDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", GetProductByIDHandler)
	getProductHandler := kithttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductHandler)

	addProductHandler := kithttp.NewServer(makeAddProductEndPoint(s), addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(makeUpdateProductEndPoint(s), updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	return r
}
func getProductByIDRequesDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productid, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productid,
	}, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil

}

func addProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)

	}

	return request, nil
}

func updateProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)

	}

	return request, nil
}
