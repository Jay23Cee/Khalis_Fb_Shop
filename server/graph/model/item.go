package model

import "time"

type Item struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"-"`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ProfileLink string `json:"profile_link"`
	ProfilePic  string `json:"profile_pic"`
}

type Photo struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"created_time"`
	From        User      `json:"from"`
	Name        string    `json:"name"`
	Picture     string    `json:"picture"`
	Link        string    `json:"link"`
}

type Comment struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"created_time"`
	From        User      `json:"from"`
	Message     string    `json:"message"`
}
