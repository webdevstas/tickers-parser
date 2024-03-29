package updater

import (
	"tickers-parser/internal/entities"
	http_client "tickers-parser/internal/services/http-client"
	"tickers-parser/pkg/utils"
)

type CryptorankCoin struct {
	Rank   int    `json:"rank,omitempty"`
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

type response struct {
	Data []CryptorankCoin `json:"data"`
}

func ParseCoinsFromCryptorank(http *http_client.HttpClient) []entities.Coin {
	var response response
	url := "https://api.cryptorank.io/v0/coins"
	http.FetchJson(url, &response)
	coins := response.Data

	return utils.Map(coins, func(el CryptorankCoin) entities.Coin {
		return CryptorankCoinToEntity(el)
	})
}

func CryptorankCoinToEntity(coin CryptorankCoin) entities.Coin {
	return entities.Coin{
		Name:    coin.Name,
		Key:     coin.Key,
		Symbol:  coin.Symbol,
		Enabled: true,
	}
}
