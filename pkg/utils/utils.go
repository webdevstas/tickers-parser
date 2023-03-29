package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tickers-parser/internal/entities"
)

func FetchJson[T any](url string, link T) error {
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != 200 {
		err = fmt.Errorf("error to get %v. Code %d. Error: %v", url, resp.StatusCode, err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		err = fmt.Errorf("error to parse body %v. %v", url, err)
		return err
	}

	err = json.Unmarshal(body, &link)

	if err != nil {
		err = fmt.Errorf("unmarshal error %v. %v", url, err)
		return err
	}

	return nil
}

func Map[F, T any](iterable []F, cb func(el F) T) []T {
	var res = make([]T, 0, len(iterable))

	for _, el := range iterable {
		res = append(res, cb(el))
	}

	return res
}

func Filter[T comparable](iterable []T, cb func(el T) bool) []T {
	var res = make([]T, 0, len(iterable))

	for _, el := range iterable {
		if cb(el) {
			res = append(res, el)
		}
	}

	return res
}

func Reduce[T, K any](iterable []T, cb func(K, T) K, initVal K) K {
	res := initVal

	for _, el := range iterable {
		res = cb(res, el)
	}

	return res
}

func RawTickerToEntity(exchangeId uint, rawTicker entities.ExchangeRawTicker) entities.Ticker {
	return entities.Ticker{
		BaseSymbol:   rawTicker.BaseSymbol,
		QuoteSymbol:  rawTicker.QuoteSymbol,
		Volume:       rawTicker.Volume,
		Bid:          rawTicker.Bid,
		Ask:          rawTicker.Ask,
		Open:         rawTicker.Open,
		High:         rawTicker.High,
		Low:          rawTicker.Low,
		Change:       rawTicker.Change,
		ExchangeID:   exchangeId,
		BaseAddress:  rawTicker.BaseAddress,
		QuoteAddress: rawTicker.QuoteAddress,
		Enabled:      false,
		Last:         rawTicker.Last,
	}
}
