package costumers

import (
	"context"

	"github.com/GolangValdezApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getCostumerListRequest struct {
	limit  int
	offset int
}

func makeGetCostumerListEndPoint(s Service) endpoint.Endpoint {

	customerList := func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(getCostumerListRequest)
		result, err := s.getCostumerList(&req)
		helper.Catch(err)
		return result, nil
	}

	return customerList
}
