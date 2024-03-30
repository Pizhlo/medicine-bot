package fsm

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/controller"
	tele "gopkg.in/telebot.v3"
)

// Дефолтное состояние бота, в котором он воспринимает любой текст как заметку
type defaultState struct {
	fsm        *FSM
	controller *controller.Controller
	name       stateName
	next       state
}

func newDefaultState(controller *controller.Controller, FSM *FSM) *defaultState {
	return &defaultState{fsm: FSM, controller: controller, name: defaultStateName, next: nil}
}

func (n *defaultState) Handle(ctx context.Context, telectx tele.Context) error {
	return nil
}

func (n *defaultState) Name() string {
	return string(n.name)
}

func (n *defaultState) Next() state {
	return n.fsm.defaultState
}
