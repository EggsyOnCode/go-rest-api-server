package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EggsyOnCode/go-rest-api-server/api"
	"github.com/EggsyOnCode/go-rest-api-server/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalServerError(w, err)
		return

	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDB()
	if err != nil {
		api.InternalServerError(w, err)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetCoinDetails(params.Username)

	if tokenDetails == nil {
		log.Error(err)
		api.InternalServerError(w, err)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: uint(tokenDetails.Coins),
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		api.InternalServerError(w, err)
		return
	}
}
