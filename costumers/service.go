package costumers

import (
	"github.com/GolangValdezApi/helper"
)

type Service interface {
	getCostumerList(params *getCostumerListRequest) (*CustomerLIst, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) getCostumerList(params *getCostumerListRequest) (*CustomerLIst, error) {

	clientes, err := s.repo.getClients(params)
	helper.Catch(err)

	totalClientes, err := s.repo.GetTotalClients()
	helper.Catch(err)

	return &CustomerLIst{Data: clientes, TotalRecortds: totalClientes}, nil

}
