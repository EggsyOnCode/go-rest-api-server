package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		Username: "alex",
		AuthToken: "123",
	},

	"human": {
		Username: "human",
		AuthToken: "123",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Username: "alex",
		Coins:    12,
	},

	"human": {
		Username: "human",
		Coins:    3123,
	},
}

func (d *mockDB) GetLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientdata = LoginDetails{}
	clientdata, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientdata
}

func (d *mockDB) GetCoinDetails(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDB() error {
	return nil
}
