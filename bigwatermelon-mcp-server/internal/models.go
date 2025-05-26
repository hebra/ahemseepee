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

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Zip       string  `json:"zip"`
	Country   string  `json:"country"`
}

type ResponseData struct {
	LastUpdated string   `json:"lastUpdated"`
	Business    string   `json:"business"`
	Location    Location `json:"location"`
	Offers      []Offer  `json:"offers"`
}

type MCPRequest struct {
	Action     string          `json:"action"`
	Parameters json.RawMessage `json:"parameters"`
	RequestID  string          `json:"request_id"`
}

type MCPResponse struct {
	Status    string      `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	RequestID string      `json:"request_id"`
}
