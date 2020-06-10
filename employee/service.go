package employee

import "github.com/GolangValdezApi/helper"

type Service interface {
	GetEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	GetEmployeesByID(params *getEmployeesByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	AddEmployee(params *addEmployeeRequest) (int64, error)
	UpdateEmployee(params *UpdateEmployeeRequest) (int64, error)
	GetDelateEmployee(params *getDeleteEmployeesByIDRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetEmployees(params *getEmployeesRequest) (*EmployeeList, error) {
	employees, err := s.repo.GetEmployees(params)
	helper.Catch(err)
	totalEmployees, err := s.repo.GetTotalEmployees()
	helper.Catch(err)
	return &EmployeeList{

		Data:         employees,
		TotalRecords: totalEmployees,
	}, nil
}

func (s *service) GetEmployeesByID(params *getEmployeesByIDRequest) (*Employee, error) {
	return s.repo.GetEmployeesByID(params.EmployeeID)

}

func (s *service) GetBestEmployee() (*BestEmployee, error) {

	return s.repo.GetBestEmployee()

}

func (s *service) AddEmployee(params *addEmployeeRequest) (int64, error) {
	return s.repo.AddEmployee(params)
}

func (s *service) UpdateEmployee(params *UpdateEmployeeRequest) (int64, error) {

	return s.repo.UpdateEmployee(params)
}

func (s *service) GetDelateEmployee(params *getDeleteEmployeesByIDRequest) (int64, error) {

	return s.repo.GetDelateEmployee(params)
}
