package _interface

import (
	"io"
)

type ReadWriter interface {
	io.Closer
	GetAllBook() ([]string, error)
}
