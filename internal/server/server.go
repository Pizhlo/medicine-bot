package server

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/controller"
	"github.com/sirupsen/logrus"
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
	err := s.controller.LoadUsers(ctx)
	if err != nil {
		logrus.Fatalf("error while loading all users: %v", err)
	}

	s.setupBot(ctx)
}

// HandleError обрабатывает ошибку: устанавливает состояние в дефолтное, передает контроллеру
func (s *Server) HandleError(ctx tele.Context, err error) {
	// обрабатываем ошибку
	s.controller.HandleError(ctx, err)
}
