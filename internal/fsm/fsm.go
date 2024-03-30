package fsm

import (
	"context"
	"sync"

	"github.com/Pizhlo/medicine-bot/internal/controller"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type stateName string

// названия состояний
const (
	defaultStateName         stateName = "default"
	startStateName           stateName = "start"
	drugNameStateName        stateName = "drug_name"
	drugDescriptionStateName stateName = "drug_description"
)

// Менеджер для управления состояниями бота
type FSM struct {
	// Состояние, в котором бот принимает название лекарства
	DrugName state
	// Состояние, в котором бот принимает описание лекарства
	DescriptionState state
	// Дефолтное состояние бота, в котором он воспринимает любой текст как заметку
	defaultState state
	// Состояние, когда бот принимает описание лекарства
	DrugDescriptionState state
	// Текущее состояние бота
	current state
	mu      sync.RWMutex
}

// Интерфейс для управления состояниями бота
type state interface {
	Handle(ctx context.Context, telectx tele.Context) error
	Name() string
	Next() state
}

func NewFSM(controller *controller.Controller) *FSM {
	fsm := &FSM{mu: sync.RWMutex{}}

	fsm.defaultState = newDefaultState(controller, fsm)

	fsm.DescriptionState = newdDrugDescriptionState(controller, fsm)

	drugName := newdrugNameStateState(controller, fsm)
	fsm.DrugName = drugName

	fsm.current = fsm.defaultState

	return fsm
}

// SetState устанавливает текущее состояние в переданное
func (f *FSM) SetState(state state) {
	f.mu.Lock()
	defer f.mu.Unlock()

	logrus.Debugf("Setting state to: %s", state.Name())

	f.current = state
}

// SetToDefault устанавливает текущее состояние FSM в дефолтное
func (f *FSM) SetToDefault() {
	f.SetState(f.defaultState)
}

func (f *FSM) Handle(ctx context.Context, telectx tele.Context) error {
	return f.current.Handle(ctx, telectx)
}

// Name возвращает название текущего состояния
func (f *FSM) Name() string {
	return f.current.Name()
}

// SetNext переключает состояние бота на следующее
func (f *FSM) SetNext() {
	f.SetState(f.current.Next())
}

// Current возвращает текущее состояние
func (f *FSM) Current() state {
	return f.current
}
