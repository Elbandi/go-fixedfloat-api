package fixedfloat

import (
	"encoding/json"
)

type jsonResponse struct {
	Code    Integer             `json:"code"`
	Message string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}
