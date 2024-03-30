package fsm

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/controller"
	tele "gopkg.in/telebot.v3"
)

type drugNameState struct {
	controller *controller.Controller
	fsm        *FSM
	name       stateName
	next       state
}

func newdrugNameStateState(controller *controller.Controller, FSM *FSM) *drugNameState {
	return &drugNameState{controller, FSM, drugNameStateName, FSM.DescriptionState}
}

func (n *drugNameState) Handle(ctx context.Context, telectx tele.Context) error {
	err := n.controller.DrugName(ctx, telectx)
	if err != nil {
		return err
	}

	n.fsm.SetNext()

	return nil
}

func (n *drugNameState) Name() string {
	return string(n.name)
}

func (n *drugNameState) Next() state {
	if n.next != nil {
		return n.next
	}
	return n.fsm.defaultState
}
