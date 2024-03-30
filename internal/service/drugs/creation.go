package drugs

import (
	"errors"
	"fmt"

	"github.com/Pizhlo/medicine-bot/internal/model"
)

// SaveName сохраняет название лекарства при создании
func (n *DrugsSrv) SaveName(userID int64, name string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.drugsMap[userID] = &model.Drug{
		UserID: userID,
		Name:   name,
	}
}

// SaveType сохраняет описание лекарства
func (n *DrugsSrv) SaveDescription(userID int64, description string) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	r, ok := n.drugsMap[userID]
	if !ok {
		return fmt.Errorf("error while getting drug by user ID: drug not found")
	}

	r.Description.String = description

	n.drugsMap[userID] = r

	return nil
}

// GetFromMemory достает из кэша лекарство в текущем состоянии (могут быть не заполнены все поля)
func (n *DrugsSrv) GetFromMemory(userID int64) (*model.Drug, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	d, ok := n.drugsMap[userID]
	if !ok {
		return nil, fmt.Errorf("error while getting drug by user ID: drug not found")
	}

	return d, nil
}

// checkFields проверяет, заполнены ли все поля в напоминании
func (n *DrugsSrv) checkFields(d *model.Drug) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if d.UserID == 0 {
		return errors.New("field UserID is not filled")
	}

	if d.Name == "" {
		return errors.New("field Name is not filled")
	}

	return nil
}
