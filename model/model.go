package model

import (
	_"fmt"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int64    `json:"quantity"`
}

type BookData struct {
	Data []*Book
}
