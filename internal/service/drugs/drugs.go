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
	drugsMap    map[int64]*model.Drug
}

type drugsEditor interface {
	Save(ctx context.Context, drug *model.Drug) error
	GetByUser(ctx context.Context, tgId int64) ([]model.Drug, error)
}

func New(drugsEditor drugsEditor) *DrugsSrv {
	return &DrugsSrv{drugsEditor: drugsEditor,
		mu:       sync.Mutex{},
		views:    make(map[int64]*view.DrugsView),
		drugsMap: make(map[int64]*model.Drug)}
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
