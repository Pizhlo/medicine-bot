package server

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

const startCmd = "/start"

func (s *Server) setupBot(ctx context.Context) {
	s.bot.Use(s.CheckUser(ctx))

	s.bot.Handle(startCmd, func(telectx tele.Context) error {
		return s.controller.StartCmd(ctx, telectx)
	})
}
