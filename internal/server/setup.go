package server

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/view"
	tele "gopkg.in/telebot.v3"
)

const startCmd = "/start"

func (s *Server) setupBot(ctx context.Context) {
	s.bot.Use(s.CheckUser(ctx))

	s.bot.Handle(startCmd, func(telectx tele.Context) error {
		err := s.controller.StartCmd(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})

	s.bot.Handle(&view.BtnBackToMenu, func(telectx tele.Context) error {
		err := s.controller.StartCmd(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})

	s.bot.Handle(&view.BtnDrugs, func(telectx tele.Context) error {
		err := s.controller.DrugsBtn(ctx, telectx)
		if err != nil {
			s.HandleError(telectx, err)
			return err
		}

		return nil
	})
}
