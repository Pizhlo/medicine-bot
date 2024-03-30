package drugs

import (
	"context"
	"sync"

	"github.com/Pizhlo/medicine-bot/internal/model"
	"github.com/Pizhlo/medicine-bot/internal/view"
)

type DrugsSrv struct {
	mu          sync.Mutex
	drugsEditor drugsEditor
	views       map[int64]*view.DrugsView
}

type drugsEditor interface {
	Save(ctx context.Context, tgID int64, drug model.Drug) error
	GetbyUser(ctx context.Context, tgId int64) ([]model.Drug, error)
}

func New(drugsEditor drugsEditor) *DrugsSrv {
	return &DrugsSrv{drugsEditor: drugsEditor,
		mu:    sync.Mutex{},
		views: make(map[int64]*view.DrugsView)}
}

func (s *DrugsSrv) SaveUsers(users []int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range users {
		s.views[u] = view.NewDrug()
	}
}

func (s *DrugsSrv) SaveUser(tgId int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.views[tgId] = view.NewDrug()
}
