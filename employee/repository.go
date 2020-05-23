package employee

import (
	"database/sql"

	"github.com/GolangValdezApi/helper"
)

//Repository interfaz almacena metodos
type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeesByID(EmployeeID int) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
}

type repository struct {
	db *sql.DB
}

//NewRepository crea la conexion entre la interfazy estructura
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {

	const sql = `SELECT 
				 id,first_name,last_name,company,email_address,job_title,business_phone,COALESCE(mobile_phone,''),fax_number,address 
				 FROM employees
				 LIMIT ? OFFSET ?`

	result, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)
	var employees []*Employee
	for result.Next() {
		employee := &Employee{}
		err = result.Scan(&employee.ID, &employee.FIrtsName, &employee.LastName, &employee.Company, &employee.EmailAddress, &employee.JobTittle, &employee.BusinessPhone, &employee.MobilePhone, &employee.FaxNumber, &employee.Address)
		helper.Catch(err)

		employees = append(employees, employee)
	}
	return employees, nil

}

func (repo *repository) GetTotalEmployees() (int64, error) {
	const sql = `SELECT COUNT(*) FROM employees`
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)

	helper.Catch(err)
	return total, nil
}

func (repo *repository) GetEmployeesByID(EmpleadoID int) (*Employee, error) {

	const sql = `SELECT id, first_name,last_name,company,email_address,job_title,business_phone,COALESCE(mobile_phone,''),fax_number,address 
	FROM employees WHERE id = ? `
	row := repo.db.QueryRow(sql, EmpleadoID)

	empleado := &Employee{}
	err := row.Scan(&empleado.ID, &empleado.FIrtsName, &empleado.LastName, &empleado.Company, &empleado.EmailAddress, &empleado.JobTittle, &empleado.BusinessPhone, &empleado.MobilePhone, &empleado.FaxNumber, &empleado.Address)

	helper.Catch(err)

	return empleado, nil

}

func (repo *repository) GetBestEmployee() (*BestEmployee, error) {
	const sql = `SELECT e.id, count(e.id) as totalventas, e.first_name, e.last_name
			FROM orders o
			INNER JOIN employees e ON o.employee_id = e.id
			GROUP BY o.employee_id
			ORDER BY totalventas desc
			limit 1`
	row := repo.db.QueryRow(sql)

	bestEmployee := &BestEmployee{}
	err := row.Scan(&bestEmployee.ID, &bestEmployee.TotalVentas, &bestEmployee.LastName, &bestEmployee.FIrtsName)
	helper.Catch(err)

	return bestEmployee, nil

}
