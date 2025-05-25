package internal

import (
	"encoding/json"
)

type Offer struct {
	ProductName string      `json:"productName"`
	Price       json.Number `json:"price"`
	Currency    string      `json:"currency"`
	Size        string      `json:"size"`
}
