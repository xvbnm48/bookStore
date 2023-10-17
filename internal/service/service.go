package service

import "github.com/xvbnm48/bookStore/internal/model"

type BookStoreService interface {
	GetAllBook() (model.GetAllBookResponse, error)
}
