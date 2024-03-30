package view

import (
	"fmt"

	"github.com/Pizhlo/medicine-bot/internal/model"
	tele "gopkg.in/telebot.v3"
)

type DrugsView struct {
	pages       []string
	currentPage int
}

func NewDrug() *DrugsView {
	return &DrugsView{pages: make([]string, 0), currentPage: 0}
}

const (
	createdFieldFormat = "02.01.2006 15:04:05"
)

var (
	// inline кнопка для переключения на предыдущую страницу
	BtnPrevPgDrugs = tele.Btn{Text: "<", Unique: "prev_pg_drugs"}
	// inline кнопка для переключения на следующую страницу
	BtnNextPgDrugs = tele.Btn{Text: ">", Unique: "next_pg_drugs"}

	// inline кнопка для переключения на первую страницу
	BtnFirstPgDrugs = tele.Btn{Text: "<<", Unique: "start_pg_drugs"}
	// inline кнопка для переключения на последнюю страницу
	BtnLastPgDrugs = tele.Btn{Text: ">>", Unique: "end_pg_drugs"}

	// inline кнопка чтобы отметить прием лекарства
	BtnAddTakeDrug = tele.Btn{Text: "Добавить прием", Unique: "add_take"}
)

// Message формирует список сообщений из моделей лекарств и возвращает первую страницу
func (v *DrugsView) Message(drugs []model.Drug) (string, error) {
	// if len(drugs) == 0 {
	// 	return drugs.UserDoesntHaveNotesMessage, nil
	// }

	v.pages = make([]string, 0)

	for _, drug := range drugs {
		var description string

		if len(drug.Description.String) == 0 {
			description = "нет"
		} else {
			description = drug.Description.String
		}

		var todayCount string
		if drug.TakeToday {
			todayCount = fmt.Sprintf("%d раз(а)", drug.TodayCount)
		} else {
			todayCount = "нет"
		}

		res := fmt.Sprintf("<b>%d. %s</b>\n\nОписание: %s\nПринимал(а) сегодня: %s\nСоздано: %s",
			drug.ViewID, drug.Name, description, todayCount, drug.Created.Format(createdFieldFormat))

		v.pages = append(v.pages, res)

	}

	v.currentPage = 0

	return v.pages[0], nil
}

// current возвращает номер текущей страницы
func (v *DrugsView) current() int {
	return v.currentPage + 1
}

// total возвращает общее количество страниц
func (v *DrugsView) total() int {
	return len(v.pages)
}

// Keyboard делает клавиатуру для навигации по страницам
func (v *DrugsView) Keyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	// если страниц 1, клавиатура не нужна
	if v.total() == 1 {
		menu.Inline(
			menu.Row(BtnAddDrug),
			menu.Row(BtnAddTakeDrug),
			menu.Row(BtnBackToMenu),
		)
		return menu
	}

	text := fmt.Sprintf("%d / %d", v.current(), v.total())

	btn := menu.Data(text, "s")

	menu.Inline(
		menu.Row(BtnFirstPgDrugs, BtnPrevPgDrugs, btn, BtnNextPgDrugs, BtnLastPgDrugs),
		menu.Row(BtnAddDrug),
		menu.Row(BtnAddTakeDrug),
		menu.Row(BtnBackToMenu),
	)

	return menu
}
