package product

import "github.com/GolangValdezApi/helper"

//Service interface
type Service interface {
	GetProductByID(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductsList, error)
	InsertProducts(params *getAddProductRequest) (int64, error)
	UpdateProducts(params *updateProductRequest) (int64, error)
	DeleteProducts(params *deleteProductRequest) (int64, error)
	GetBestSellers() (*ProductTopResponse, error)
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

	helper.Catch(err)
	totalProducts, err := s.repo.GetTotalProducts()
	helper.Catch(err)

	return &ProductsList{Data: product, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProducts(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProducts(params)
}

func (s *service) UpdateProducts(params *updateProductRequest) (int64, error) {
	return s.repo.UpdateProducts(params)
}

func (s *service) DeleteProducts(params *deleteProductRequest) (int64, error) {

	return s.repo.DeleteProducts(params)
}

func (s *service) GetBestSellers() (*ProductTopResponse, error) {
	products, err := s.repo.getBestSellers()
	helper.Catch(err)
	totalVentas, err := s.repo.getTotalVentas()
	helper.Catch(err)
	return &ProductTopResponse{Data: products, TotalVentas: totalVentas}, err
}
