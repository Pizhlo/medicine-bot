package drugs

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

func (s *DrugsSrv) GetByUserID(ctx context.Context, tgID int64) (string, *tele.ReplyMarkup, error) {
	drugs, err := s.drugsEditor.GetByUser(ctx, tgID)
	if err != nil {
		return "", nil, err
	}

	msg, err := s.views[tgID].Message(drugs)
	if err != nil {
		return "", nil, err
	}

	return msg, s.views[tgID].Keyboard(), nil
}

// NextPage обрабатывает кнопку переключения на следующую страницу
func (s *DrugsSrv) NextPage(userID int64) (string, *tele.ReplyMarkup) {
	return s.views[userID].Next(), s.views[userID].Keyboard()
}

// PrevPage обрабатывает кнопку переключения на предыдущую страницу
func (s *DrugsSrv) PrevPage(userID int64) (string, *tele.ReplyMarkup) {
	return s.views[userID].Previous(), s.views[userID].Keyboard()
}

// LastPage обрабатывает кнопку переключения на последнюю страницу
func (s *DrugsSrv) LastPage(userID int64) (string, *tele.ReplyMarkup) {
	return s.views[userID].Last(), s.views[userID].Keyboard()
}

// FirstPage обрабатывает кнопку переключения на первую страницу
func (s *DrugsSrv) FirstPage(userID int64) (string, *tele.ReplyMarkup) {
	return s.views[userID].First(), s.views[userID].Keyboard()
}
