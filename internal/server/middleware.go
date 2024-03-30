package server

import (
	"context"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

// CheckUser проверяет, зарегистрирован ли пользователь. Если нет - запрашивает геолокацию.
// Если да - обрабатывает запрос
func (s *Server) CheckUser(contxt context.Context) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			if !s.controller.CheckUser(contxt, ctx.Chat().ID) {
				logrus.Debugf("user %d not found in map. Saving...", ctx.Chat().ID)
				err := s.SaveUser(contxt, ctx)
				if err != nil {
					logrus.Errorf("error while saving user: %v", err)
				}
			}
			return next(ctx)
		}
	}
}
