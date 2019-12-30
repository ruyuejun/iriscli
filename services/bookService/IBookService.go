package bookService

import (
	"iriscli/datamodels"
)
type IBookService interface {
	GetBookList (m map[string]interface{}) (result datamodels.Result)
}

func NewBookService() IBookService{
	return &bookService{}
}