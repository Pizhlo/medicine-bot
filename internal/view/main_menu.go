package view

import tele "gopkg.in/telebot.v3"

var (
	// -------------- главное меню --------------

	BtnDrugs = tele.Btn{Text: "💊Мои лекарства", Unique: "drugs"}
)

// MainMenu возвращает главное меню
func MainMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	menu.Inline(
		menu.Row(BtnDrugs),
	)

	return menu
}
