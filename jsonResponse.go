package fixedfloat

import (
	"encoding/json"
)

type jsonResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}
