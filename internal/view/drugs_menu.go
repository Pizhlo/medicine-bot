package view

import tele "gopkg.in/telebot.v3"

var (
	// inline кнопка добавить лекарство
	BtnAddDrug = tele.Btn{Text: "💊Добавить лекарство", Unique: "add_drug"}

	// inline кнопка чтобы пропустить заполнение описания лекарства
	BtnSkipDescription = tele.Btn{Text: "Пропустить", Unique: "skip_description"}
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

// SkipDescriptionMenu возвращает меню с кнопками: пропустить описание, назад в меню
func SkipDescriptionMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	menu.Inline(
		menu.Row(BtnSkipDescription),
		menu.Row(BtnBackToMenu),
	)

	return menu
}
