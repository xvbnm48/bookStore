package model

type Book struct {
	ID       string
	Title    string
	Sinopsis string
	ISBN     string
}
type GetAllBookResponse struct {
	Books []Book
}
