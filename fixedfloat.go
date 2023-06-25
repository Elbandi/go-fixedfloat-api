package fixedfloat

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

const (
	ApiBase = "https://ff.io/api/v2" // FixedFloat API endpoint
)

// New returns an instantiated FixedFloat struct
func New(apiKey, apiSecret string) *FixedFloat {
	client := NewClient(apiKey, apiSecret)
	return &FixedFloat{client}
}

// NewWithCustomHttpClient returns an instantiated FixedFloat struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *FixedFloat {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &FixedFloat{client}
}

// NewWithCustomTimeout returns an instantiated FixedFloat struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *FixedFloat {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &FixedFloat{client}
}

// handleErr gets JSON response from FixedFloat API en deal with error
func handleErr(r jsonResponse) error {
	if r.Code != 0 {
		return errors.New(r.Message)
	}
	return nil
}

// FixedFloat represent a FixedFloat client
type FixedFloat struct {
	client *client
}

// SetDebug set enable/disable http request/response dump
func (ff *FixedFloat) SetDebug(enable bool) {
	ff.client.debug = enable
}

type Currency struct {
	Code     string `json:"code"`
	Coin     string `json:"coin"`
	Network  string `json:"network"`
	Name     string `json:"name"`
	Recv     Bool   `json:"recv"`
	Send     Bool   `json:"send"`
	Tag      string `json:"tag"`
	Priority uint   `json:"priority"`
}

// GetCurrencies getting a list of currencies supported by the FixedFloat service
func (ff *FixedFloat) GetCurrencies() (currencies []Currency, err error) {
	r, err := ff.client.do("POST", "ccies", nil, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Data, &currencies)
	return
}

type Rate struct {
	Code      string  `json:"code"`
	Coin      string  `json:"coin"`
	Network   string  `json:"network"`
	Amount    float64 `json:"amount,string"`
	Rate      float64 `json:"rate,string"`
	Precision uint8   `json:"precision"`
	Min       float64 `json:"min,string"`
	Max       float64 `json:"max,string"`
	Usd       float64 `json:"usd,string"`
	//	Btc     float64 `json:"btc,string"`
}

// GetRate getting the exchange rate of a pair of currencies in the selected direction and type of rate.
func (ff *FixedFloat) GetRate(fromCurrency string, toCurrency string, amount float64) (Rate, Rate, error) {
	payload := map[string]interface{}{
		"fromCcy":   fromCurrency,
		"toCcy":     toCurrency,
		"amount":    amount,
		"direction": "from",
		"type":      "float",
	}
	r, err := ff.client.do("POST", "price", payload, true)
	if err != nil {
		return Rate{}, Rate{}, err
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return Rate{}, Rate{}, err
	}
	if err = handleErr(response); err != nil {
		return Rate{}, Rate{}, err
	}
	var rates struct {
		From  Rate     `json:"from"`
		To    Rate     `json:"to"`
		Error []string `json:"errors"`
	}
	if err = json.Unmarshal(response.Data, &rates); err != nil {
		return Rate{}, Rate{}, err
	}
	if len(rates.Error) > 0 {
		return Rate{}, Rate{}, errors.New(strings.Join(rates.Error, ","))
	}
	return rates.From, rates.To, nil
}
