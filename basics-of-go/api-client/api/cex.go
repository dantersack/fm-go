package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dantejs.com/go/api-client/model"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*model.Rate, error) {
	if len(currency) == 0 {
		return nil, fmt.Errorf("please provide a valid currency")
	}
	var cexResponse CEXResponse
	res, err := http.Get(fmt.Sprintf(apiUrl, currency))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &cexResponse)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := model.Rate{Currency: currency, Price: cexResponse.Bid}
	return &rate, nil
}
