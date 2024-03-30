package server

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/messages"
	"github.com/Pizhlo/medicine-bot/internal/view"
	tele "gopkg.in/telebot.v3"
)

const startCmd = "/start"

func (s *Server) setupBot(ctx context.Context) {
	s.bot.Use(s.CheckUser(ctx))

	s.bot.Handle(tele.OnText, func(telectx tele.Context) error {
		err := s.fsm[telectx.Chat().ID].Handle(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})

	// старт
	s.bot.Handle(startCmd, func(telectx tele.Context) error {
		err := s.controller.StartCmd(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})

	// назад в меню
	s.bot.Handle(&view.BtnBackToMenu, func(telectx tele.Context) error {
		err := s.controller.StartCmd(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})

	// просмотр лекарств
	s.bot.Handle(&view.BtnDrugs, func(telectx tele.Context) error {
		err := s.controller.DrugsBtn(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})

	// навигация по страницам

	// предыдущая страница
	s.bot.Handle(&view.BtnPrevPgDrugs, func(c tele.Context) error {
		err := s.controller.PrevPageReminders(ctx, c)
		if err != nil {
			s.HandleError(c, err)
			return err
		}

		return nil
	})

	// следующая страница
	s.bot.Handle(&view.BtnNextPgDrugs, func(c tele.Context) error {
		err := s.controller.NextPageReminders(ctx, c)
		if err != nil {
			s.HandleError(c, err)
			return err
		}

		return nil
	})

	// первая страница
	s.bot.Handle(&view.BtnFirstPgDrugs, func(c tele.Context) error {
		err := s.controller.FirstPageReminders(ctx, c)
		if err != nil {
			s.HandleError(c, err)
			return err
		}

		return nil
	})

	// последняя страница
	s.bot.Handle(&view.BtnLastPgDrugs, func(c tele.Context) error {
		err := s.controller.LastPageReminders(ctx, c)
		if err != nil {
			s.HandleError(c, err)
			return err
		}

		return nil
	})

	// добавить лекарство
	s.bot.Handle(&view.BtnAddDrug, func(telectx tele.Context) error {
		s.fsm[telectx.Chat().ID].SetState(s.fsm[telectx.Chat().ID].DrugName)

		err := telectx.EditOrSend(messages.DrugNameMessage, view.BackToMenuBtn())
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})

	// пропустить описание
	s.bot.Handle(&view.BtnSkipDescription, func(telectx tele.Context) error {
		s.fsm[telectx.Chat().ID].SetToDefault()

		err := s.controller.DrugDescription(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})
}
