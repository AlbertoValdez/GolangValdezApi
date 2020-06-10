package employee

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/GolangValdezApi/helper"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {

	r := chi.NewRouter()
	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndPoint(s), getEmployesRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginado", getEmployeesHandler)

	getEmployeeByIDHandler := kithttp.NewServer(makeGetEmployeeByIDEndPoint(s), getEmployeeByIDRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIDHandler)

	getBestEmployeeHanlder := kithttp.NewServer(makeGetBestEmployeeEndPoint(s), getBestEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/mejorempleado", getBestEmployeeHanlder)

	getAddEmployeeHandler := kithttp.NewServer(makeAddEmployeeEndPoint(s), getAddEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", getAddEmployeeHandler)

	getUpdateEmployeeHandler := kithttp.NewServer(makeUpdateEmployeEndPoint(s), getUpdateRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", getUpdateEmployeeHandler)

	getDeleteEmploteeHandler := kithttp.NewServer(makeDeleteEmployeEndPoint(s), getDeleteRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", getDeleteEmploteeHandler)

	return r

}

func getEmployesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	helper.Catch(err)
	return request, nil

}

func getEmployeeByIDRequestDecoder(Context context.Context, r *http.Request) (interface{}, error) {

	EmpleadoID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getEmployeesByIDRequest{
		EmployeeID: EmpleadoID,
	}, nil

}

func getBestEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {

	return getBestEmployeeRequest{}, nil
}

func getAddEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)

	return request, nil
}
func getUpdateRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {

	request := UpdateEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getDeleteRequestDecoder(Context context.Context, r *http.Request) (interface{}, error) {

	EmpleadoID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getDeleteEmployeesByIDRequest{
		EmployeeID: EmpleadoID,
	}, nil
}
