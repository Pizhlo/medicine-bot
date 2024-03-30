package fsm

import (
	"context"

	"github.com/Pizhlo/medicine-bot/internal/controller"
	tele "gopkg.in/telebot.v3"
)

type drugDescriptionState struct {
	controller *controller.Controller
	fsm        *FSM
	name       stateName
	next       state
}

func newdDrugDescriptionState(controller *controller.Controller, FSM *FSM) *drugDescriptionState {
	return &drugDescriptionState{controller, FSM, drugDescriptionStateName, nil}
}

func (n *drugDescriptionState) Handle(ctx context.Context, telectx tele.Context) error {
	err := n.controller.DrugDescription(ctx, telectx)
	if err != nil {
		return err
	}

	n.fsm.SetNext()

	return nil
}

func (n *drugDescriptionState) Name() string {
	return string(n.name)
}

func (n *drugDescriptionState) Next() state {
	if n.next != nil {
		return n.next
	}
	return n.fsm.defaultState
}
