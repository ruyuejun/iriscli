package bookRepo

import (
	"iriscli/datamodels"
	"iriscli/datasource"
)

var db = datasource.GetDB()

type bookRepository struct {}

func (n bookRepository) GetBookList(m map[string]interface{})(total int,books []datamodels.Book){

	db.Table("book").Count(&total)

	db.Find(&books)

	//err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"])-1)*cast.ToInt(m["size"])).Find(&books).Error
	//if err!=nil {
	//	panic("select Error")
	//}

	return
}
