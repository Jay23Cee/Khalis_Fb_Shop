// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model


type ItemInput struct {
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	ID           int     `json:"id"`
	CustomerName string  `json:"customerName"`
	OrderAmount  float64 `json:"orderAmount"`
	Items        []*Item `json:"items"`
}

type OrderInput struct {
	CustomerName string       `json:"customerName"`
	OrderAmount  float64      `json:"orderAmount"`
	Items        []*ItemInput `json:"items"`
}
