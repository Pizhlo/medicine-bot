package controller

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/messages"
	"github.com/Pizhlo/medicine-bot/internal/view"
	tele "gopkg.in/telebot.v3"
)

func (c *Controller) StartCmd(ctx context.Context, telectx tele.Context) error {
	return telectx.EditOrSend(messages.StartMessage, &tele.SendOptions{
		ReplyMarkup: view.MainMenu(),
	})
}
