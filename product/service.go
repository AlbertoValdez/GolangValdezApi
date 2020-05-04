package product

//Service interface
type Service interface {
	GetProductByID(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductsList, error)
	InsertProducts(params *getAddProductRequest) (int64, error)
}

type service struct {
	repo Repository
}

//Ns new services
func Ns(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) GetProductByID(param *getProductByIDRequest) (*Product, error) {
	return s.repo.GetProductByID(param.ProductID)
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductsList, error) {
	product, err := s.repo.GetProducts(params)

	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}

	return &ProductsList{Data: product, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProducts(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProducts(params)
}
