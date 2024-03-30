package controller

import (
	"context"
	"errors"

	api_errors "github.com/Pizhlo/medicine-bot/internal/errors"
	"github.com/Pizhlo/medicine-bot/internal/messages"
	"github.com/Pizhlo/medicine-bot/internal/view"
	tele "gopkg.in/telebot.v3"
)

// DrugsBtn обрабатывает кнопку "Мои лекарства"
func (c *Controller) DrugsBtn(ctx context.Context, telectx tele.Context) error {
	msg, kb, err := c.drugSrv.GetByUserID(ctx, telectx.Chat().ID)
	if err != nil {
		if errors.Is(err, api_errors.ErrDrugsNotFound) {
			return telectx.EditOrSend(messages.DrugsNotFoundMessage, view.AddDrugBtn())
		}
	}

	return telectx.EditOrSend(msg, &tele.SendOptions{
		ParseMode:   htmlParseMode,
		ReplyMarkup: kb,
	})
}
