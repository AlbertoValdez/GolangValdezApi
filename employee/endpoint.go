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

type addEmployeeRequest struct {
	LastName      string
	FIrtsName     string
	Company       string
	EmailAddress  string
	JobTittle     string
	BusinessPhone string
	MobilePhone   string
	FaxNumber     string
	Address       string
}

type UpdateEmployeeRequest struct {
	ID            int64
	LastName      string
	FIrtsName     string
	Company       string
	EmailAddress  string
	JobTittle     string
	BusinessPhone string
	MobilePhone   string
	FaxNumber     string
	HomePhone     string
	Address       string
}

type getDeleteEmployeesByIDRequest struct {
	EmployeeID int
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

	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {

		result, err := s.GetBestEmployee()
		helper.Catch(err)

		return result, nil
	}

	return getBestEmployeeEndPoint
}

func makeAddEmployeeEndPoint(s Service) endpoint.Endpoint {

	addEmployee := func(_ context.Context, request interface{}) (interface{}, error) {

		req := request.(addEmployeeRequest)
		result, err := s.AddEmployee(&req)
		helper.Catch(err)

		return result, nil
	}

	return addEmployee

}

func makeUpdateEmployeEndPoint(s Service) endpoint.Endpoint {

	updateEmployee := func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEmployeeRequest)
		result, err := s.UpdateEmployee(&req)
		helper.Catch(err)
		return result, nil
	}
	return updateEmployee

}

func makeDeleteEmployeEndPoint(s Service) endpoint.Endpoint {

	deleteEmployee := func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getDeleteEmployeesByIDRequest)
		result, err := s.GetDelateEmployee(&req)
		helper.Catch(err)
		return result, nil

	}

	return deleteEmployee

}
