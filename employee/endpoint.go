package employee

import (
	"context"

	"github.com/GolangValdezApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

type getEmployeesByIDRequest struct {
	EmployeeID int
}

type getBestEmployeeRequest struct {
}

func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployeesEndPoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getEmployeesRequest)
		result, err := s.GetEmployees(&req)
		helper.Catch(err)
		return result, nil
	}

	return getEmployeesEndPoint
}

func makeGetEmployeeByIDEndPoint(s Service) endpoint.Endpoint {

	getEmployeesByID := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getEmployeesByIDRequest)
		result, err := s.GetEmployeesByID(&req)
		helper.Catch(err)
		return result, nil

	}

	return getEmployeesByID

}

func makeGetBestEmployeeEndPoint(s Service) endpoint.Endpoint {

	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (_ interface{}, err error) {

		result, err := s.GetBestEmployee()
		helper.Catch(err)

		return result, nil
	}

	return getBestEmployeeEndPoint
}
