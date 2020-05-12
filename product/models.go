package product

//Product Estructura
type Product struct {
	ID           int     `json:"id"`
	ProductCode  string  `json:"productCode"`
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	StandardCost float64 `json:"standardCost"`
	ListPrice    float64 `json:"listPrice"`
	Category     string  `json:"category"`
}

//ProductsList Estructura
type ProductsList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}

//ProductTop Estructura
type ProductTop struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	Vendidos    float64 `json:"vendidos"`
}

//ProductTopResponse Estructura
type ProductTopResponse struct {
	Data        []*ProductTop `json:"data"`
	TotalVentas float64       `json:"totalVentas"`
}
