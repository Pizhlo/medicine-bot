package controller

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/messages"
	"github.com/Pizhlo/medicine-bot/internal/service/drugs"
	"github.com/Pizhlo/medicine-bot/internal/service/user"
	"github.com/Pizhlo/medicine-bot/internal/view"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type Controller struct {
	userSrv *user.UserSrv
	drugSrv *drugs.DrugsSrv
}

const (
	htmlParseMode = "HTML"
)

func New(userSrv *user.UserSrv, drugSrv *drugs.DrugsSrv) *Controller {
	return &Controller{userSrv: userSrv, drugSrv: drugSrv}
}

// CheckUser проверяет, известен ли пользователь боту
func (c *Controller) CheckUser(ctx context.Context, tgID int64) bool {
	return c.userSrv.CheckUser(ctx, tgID)
}

func (c *Controller) SaveUser(ctx context.Context, tgID int64) error {
	c.drugSrv.SaveUser(tgID)
	return c.userSrv.SaveUser(ctx, tgID)
}

func (c *Controller) LoadUsers(ctx context.Context) error {
	users, err := c.userSrv.GetAllUsers(ctx)
	if err != nil {
		return nil
	}

	c.drugSrv.SaveUsers(users)

	return c.userSrv.LoadUsers(ctx, users)
}

// HandleError сообщает об ошибке в канал.
// Также сообщает пользователю об ошибке - редактирует сообщение
func (c *Controller) HandleError(ctx tele.Context, err error) {
	editErr := ctx.EditOrSend(messages.ErrorMessageUser, view.BackToMenuBtn())
	if editErr != nil {
		logrus.Errorf("Error while sending error message to user. Error: %+v\n", editErr)
	}
}
