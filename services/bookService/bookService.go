package bookService

import (
	"iriscli/datamodels"
	"iriscli/datarepo/bookRepo"
)

type bookService struct {}

var br = bookRepo.NewBookRepo()

func (bs bookService)GetBookList (m map[string]interface{}) (result datamodels.Result){
	total,books := br.GetBookList(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = books
	result.Data = maps
	result.Code = 0
	result.Msg ="SUCCESS"
	return
}
