package services

import (
	"clima/src/types"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	unexpectedError = errors.New("ocorreu um erro inesperado")
)

func Fetch(url string) (types.Weather, error) {
	var weather types.Weather

	res, err := http.Get(url)
	if err != nil {
		return weather, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return weather, unexpectedError
	}

	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		return weather, err
	}

	return weather, nil
}
