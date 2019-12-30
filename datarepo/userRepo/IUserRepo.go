package userRepo

import "iriscli/datamodels"

type IUserRepo interface {
	GetUserByUserNameAndPwd(username string, password string) (user datamodels.User)
	GetUserByUsername(username string) (user datamodels.User)
	Save(user datamodels.User) (int, datamodels.User)
}

func NewUserRepo() IUserRepo {
	return &userRepository{}
}