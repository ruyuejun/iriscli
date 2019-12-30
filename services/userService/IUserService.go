package userService

import "iriscli/datamodels"

type IUser interface {
	Login(m map[string]string) (result datamodels.Result)
	Save(user datamodels.User) (result datamodels.Result)
}
