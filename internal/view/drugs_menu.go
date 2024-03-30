package view

import tele "gopkg.in/telebot.v3"

var (
	// inline кнопка добавить лекарство
	BtnAddDrug = tele.Btn{Text: "Добавить лекарство", Unique: "add_drug"}
)

// AddDrugBtn возвращает меню с кнопками: добавить лекарство, назад в меню
func AddDrugBtn() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	menu.Inline(
		menu.Row(BtnAddDrug),
		menu.Row(BtnBackToMenu),
	)

	return menu
}
