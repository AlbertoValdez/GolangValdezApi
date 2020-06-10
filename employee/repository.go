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
	AddEmployee(params *addEmployeeRequest) (int64, error)

	UpdateEmployee(params *UpdateEmployeeRequest) (int64, error)

	GetDelateEmployee(params *getDeleteEmployeesByIDRequest) (int64, error)
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

func (repo *repository) AddEmployee(params *addEmployeeRequest) (int64, error) {

	const sql = `INSERT INTO employees (first_name,last_name,company,email_address,job_title,business_phone,mobile_phone,fax_number,address)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?,? )`

	result, err := repo.db.Exec(sql, params.FIrtsName, params.LastName, params.Company, params.EmailAddress, params.JobTittle, params.BusinessPhone, params.MobilePhone, params.FaxNumber, params.Address)
	helper.Catch(err)

	id, err := result.LastInsertId()
	helper.Catch(err)
	return id, nil

}

func (repo *repository) UpdateEmployee(params *UpdateEmployeeRequest) (int64, error) {

	const sql = `UPDATE   employees SET 
				first_name = ?,
				last_name =  ?,
				company  =   ?,
				email_address = ?,
				job_title = ?,
				business_phone = ?,
				mobile_phone = ?,
				fax_number = ?,
				home_phone = ?, 
				address = ? 
				WHERE id = ?`

	_, err := repo.db.Exec(sql, params.FIrtsName, params.LastName, params.Company, params.EmailAddress, params.JobTittle, params.BusinessPhone, params.MobilePhone, params.FaxNumber, params.HomePhone, params.Address, params.ID)
	helper.Catch(err)
	id := params.ID
	return id, nil

}

func (repo *repository) GetDelateEmployee(params *getDeleteEmployeesByIDRequest) (int64, error) {

	const sql = `DELETE FROM employees WHERE id = ? `
	result, err := repo.db.Exec(sql, params.EmployeeID)
	helper.Catch(err)

	count, err := result.RowsAffected()
	helper.Catch(err)

	return count, nil

}
