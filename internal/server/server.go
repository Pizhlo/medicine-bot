package server

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/controller"
	tele "gopkg.in/telebot.v3"
)

type Server struct {
	bot        *tele.Bot
	controller *controller.Controller
}

func New(bot *tele.Bot, controller *controller.Controller) *Server {
	return &Server{bot: bot,
		controller: controller}
}

func (s *Server) Start(ctx context.Context) {
	s.setupBot(ctx)
}
