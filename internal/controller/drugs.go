package controller

import (
	"context"
	"errors"
	"fmt"

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

// NextPageReminders обрабатывает кнопку переключения на следующую страницу
func (c *Controller) NextPageReminders(ctx context.Context, telectx tele.Context) error {
	page, kb := c.drugSrv.NextPage(telectx.Chat().ID)

	err := telectx.EditOrSend(page, &tele.SendOptions{
		ReplyMarkup: kb,
		ParseMode:   htmlParseMode,
	})

	// если пришла ошибка о том, что сообщение не изменено - игнорируем.
	// такая ошибка происходит, если быть на первой странице и нажать кнопку "первая страница".
	// то же самое происходит и с последней страницей
	if err != nil {
		switch t := err.(type) {
		case *tele.Error:
			if t.Description == "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message (400)" {
				break
			}
		default:
			return err
		}
	}

	return nil
}

// NextPageReminders обрабатывает кнопку переключения на предыдущую страницу
func (c *Controller) PrevPageReminders(ctx context.Context, telectx tele.Context) error {
	page, kb := c.drugSrv.PrevPage(telectx.Chat().ID)

	err := telectx.EditOrSend(page, &tele.SendOptions{
		ReplyMarkup: kb,
		ParseMode:   htmlParseMode,
	})

	// если пришла ошибка о том, что сообщение не изменено - игнорируем.
	// такая ошибка происходит, если быть на первой странице и нажать кнопку "первая страница".
	// то же самое происходит и с последней страницей
	if err != nil {
		switch t := err.(type) {
		case *tele.Error:
			if t.Description == "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message (400)" {
				break
			}
		default:
			return err
		}
	}

	return nil
}

// NextPageReminders обрабатывает кнопку переключения на последнюю страницу
func (c *Controller) LastPageReminders(ctx context.Context, telectx tele.Context) error {
	page, kb := c.drugSrv.LastPage(telectx.Chat().ID)

	err := telectx.EditOrSend(page, &tele.SendOptions{
		ReplyMarkup: kb,
		ParseMode:   htmlParseMode,
	})

	// если пришла ошибка о том, что сообщение не изменено - игнорируем.
	// такая ошибка происходит, если быть на первой странице и нажать кнопку "первая страница".
	// то же самое происходит и с последней страницей
	if err != nil {
		switch t := err.(type) {
		case *tele.Error:
			if t.Description == "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message (400)" {
				break
			}
		default:
			return err
		}
	}

	return nil
}

// NextPageReminders обрабатывает кнопку переключения на первую страницу
func (c *Controller) FirstPageReminders(ctx context.Context, telectx tele.Context) error {
	page, kb := c.drugSrv.FirstPage(telectx.Chat().ID)

	err := telectx.EditOrSend(page, &tele.SendOptions{
		ReplyMarkup: kb,
		ParseMode:   htmlParseMode,
	})

	// если пришла ошибка о том, что сообщение не изменено - игнорируем.
	// такая ошибка происходит, если быть на первой странице и нажать кнопку "первая страница".
	// то же самое происходит и с последней страницей

	if err != nil {
		switch t := err.(type) {
		case *tele.Error:
			if t.Description == "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message (400)" {
				break
			}
		default:
			return err
		}
	}

	return nil
}

// DrugName сохраняет название лекарства
func (c *Controller) DrugName(ctx context.Context, telectx tele.Context) error {
	if !telectx.Message().Sender.IsBot {
		c.drugSrv.SaveName(telectx.Chat().ID, telectx.Message().Text)
	}

	return telectx.EditOrSend(messages.DrugDescriptionMessage, &tele.SendOptions{
		ParseMode:   htmlParseMode,
		ReplyMarkup: view.SkipDescriptionMenu(),
	})
}

// DrugDescription сохраняет описание лекарства
func (c *Controller) DrugDescription(ctx context.Context, telectx tele.Context) error {
	if !telectx.Message().Sender.IsBot {
		c.drugSrv.SaveDescription(telectx.Chat().ID, telectx.Message().Text)
	}

	err := c.drugSrv.SaveDrug(ctx, telectx.Chat().ID)
	if err != nil {
		return err
	}

	d, err := c.drugSrv.GetFromMemory(telectx.Chat().ID)
	if err != nil {
		return err
	}

	txt := fmt.Sprintf(messages.DescriptionSavedMessage, d.Name)

	return telectx.EditOrSend(txt, &tele.SendOptions{
		ParseMode:   htmlParseMode,
		ReplyMarkup: view.MainMenu(),
	})
}
