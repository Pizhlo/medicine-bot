package view

import tele "gopkg.in/telebot.v3"

var (
	// -------------- главное меню --------------

	BtnDrugs = tele.Btn{Text: "💊Мои лекарства", Unique: "drugs"}

	// inline кнопка для возвращения в меню
	BtnBackToMenu = tele.Btn{Text: "⬅️Меню", Unique: "menu"}
)

// MainMenu возвращает главное меню
func MainMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	menu.Inline(
		menu.Row(BtnDrugs),
	)

	return menu
}

// BackToMenuBtn возвращает кнопку возврата в меню
func BackToMenuBtn() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	menu.Inline(
		menu.Row(BtnBackToMenu),
	)

	return menu
}
