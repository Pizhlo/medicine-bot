package controller

import "github.com/Pizhlo/medicine-bot/internal/service/user"

type Controller struct {
	userSrv *user.UserSrv
}

func New(userSrv *user.UserSrv) *Controller {
	return &Controller{userSrv: userSrv}
}
