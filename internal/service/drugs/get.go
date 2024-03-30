package drugs

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

func (s *DrugsSrv) GetByUserID(ctx context.Context, tgID int64) (string, *tele.ReplyMarkup, error) {
	drugs, err := s.drugsEditor.GetbyUser(ctx, tgID)
	if err != nil {
		return "", nil, err
	}

	msg, err := s.views[tgID].Message(drugs)
	if err != nil {
		return "", nil, err
	}

	return msg, s.views[tgID].Keyboard(), nil
}
