package bookRepo

import "iriscli/datamodels"

type IBookRepo interface {
	GetBookList(m map[string]interface{})(total int,books []datamodels.Book)
}

func NewBookRepo() IBookRepo{
	return &bookRepository{}
}