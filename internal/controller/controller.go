package controller

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/service/user"
)

type Controller struct {
	userSrv *user.UserSrv
}

func New(userSrv *user.UserSrv) *Controller {
	return &Controller{userSrv: userSrv}
}

// CheckUser проверяет, известен ли пользователь боту
func (c *Controller) CheckUser(ctx context.Context, tgID int64) bool {
	return c.userSrv.CheckUser(ctx, tgID)
}

func (c *Controller) SaveUser(ctx context.Context, tgID int64) error {
	return c.userSrv.SaveUser(ctx, tgID)
}

func (c *Controller) LoadUsers(ctx context.Context) error {
	return c.userSrv.LoadUsers(ctx)
}
