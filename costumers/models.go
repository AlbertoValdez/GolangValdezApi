package costumers

type Customer struct {
	ID            int    `json:"id"`
	FirtsName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Address       string `json:"address"`
	BusniessPhone string `json:"busniessPhone"`
	City          string `json:"city"`
	Company       string `json:"company"`
}

type CustomerLIst struct {
	Data          []*Customer `json:"data"`
	TotalRecortds int64       `json:"totalRecords"`
}
