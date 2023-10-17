package server

import (
	"github.com/xvbnm48/bookStore/internal/repository"
	"github.com/xvbnm48/bookStore/internal/service"
)

type BookStoreService struct {
	repo *repository.Repo
}

// TODO: MAKE GET ALL BOOK SERVICE
func NewBookStoreServer(repo repository.Repo) service.BookStoreService {
	return &BookStoreService{
		repo: repo,
	}
}
