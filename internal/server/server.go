package server

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/controller"
	"github.com/Pizhlo/medicine-bot/internal/fsm"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type Server struct {
	bot        *tele.Bot
	controller *controller.Controller
	fsm        map[int64]*fsm.FSM
}

func New(bot *tele.Bot, controller *controller.Controller) *Server {
	return &Server{bot: bot,
		controller: controller,
		fsm:        make(map[int64]*fsm.FSM)}
}

func (s *Server) Start(ctx context.Context) {
	users, err := s.controller.GetAllUsers(ctx)
	if err != nil {
		logrus.Fatalf("error while getting all users: %v", err)
	}

	err = s.controller.LoadUsers(ctx, users)
	if err != nil {
		logrus.Fatalf("error while loading all users: %v", err)
	}

	for _, u := range users {
		s.RegisterUserInFSM(u)
	}

	s.setupBot(ctx)
}

// HandleError обрабатывает ошибку: устанавливает состояние в дефолтное, передает контроллеру
func (s *Server) HandleError(ctx tele.Context, err error) {
	// обрабатываем ошибку
	s.controller.HandleError(ctx, err)
}

func (s *Server) RegisterUserInFSM(userID int64) {
	s.fsm[userID] = fsm.NewFSM(s.controller)
}

func (s *Server) SaveUser(ctx context.Context, telectx tele.Context) error {
	s.fsm[telectx.Chat().ID] = fsm.NewFSM(s.controller)
	return s.controller.SaveUser(ctx, telectx.Chat().ID)
}
