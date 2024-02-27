package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username  string
	AuthToken string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetLoginDetails(username string) *LoginDetails
	GetCoinDetails(username string) *CoinDetails
	SetupDB() error
}

func NewDB() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var error error = database.SetupDB()
	if error != nil {
		log.Error(error)
		return nil, error
	}

	return &database, nil
}
