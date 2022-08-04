package model

// Book properties
type Book struct {
	BookID uint   `json:"book_id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Amount uint   `json:"amount"`
	State  string `json:"state"`
}

type NewBook struct {
	Name string `json:"name"`
}

type UpdateBook struct {
	State string `json:"state"`
}

type BookResponse struct {
	ID     *int   `json:"id"`
	BookID int    `json:"book_id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Amount uint   `json:"amount"`
}
