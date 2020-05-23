package employee

//Employee Estructura modelo vista para el cliente
type Employee struct {
	ID            int    `json:"id"`
	LastName      string `json:"lastName"`
	FIrtsName     string `json:"firstName"`
	Company       string `json:"company"`
	EmailAddress  string `json:"emailAddress"`
	JobTittle     string `json:"jobTitle"`
	BusinessPhone string `json:"busniessPhone"`
	MobilePhone   string `json:"mobilePhone"`
	FaxNumber     string `json:"faxNumber"`
	Address       string `json:"address"`
}

//EmployeeList Estructura
type EmployeeList struct {
	Data         []*Employee `json:"data"`
	TotalRecords int64       `json:"totalRecord"`
}

type BestEmployee struct {
	ID          int    `json:"id"`
	TotalVentas string `json:"totalVentas"`
	LastName    string `json:"lastName"`
	FIrtsName   string `json:"firstName"`
}
