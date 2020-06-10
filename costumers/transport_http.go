package costumers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GolangValdezApi/helper"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {

	r := chi.NewRouter()

	getCostumerListHandler := kithttp.NewServer(makeGetCostumerListEndPoint(s), getCostumerListRequestDecoder, kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginados", getCostumerListHandler)

	return r
}

func getCostumerListRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {

	req := getCostumerListRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)

	helper.Catch(err)
	return req, nil

}
