package api

import (
	"encoding/json"
	"net/http"
) 

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceParams struct {
	Code    int
	Balance int64
} 

type Error struct {
	Code    int
	Message string
}
