package user

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

type UserSrv struct {
	mu         sync.Mutex
	userEditor userEditor
	usersMap   map[int64]struct{}
}

type userEditor interface {
	Save(ctx context.Context, tgID int64) error
	GetAll(ctx context.Context) ([]int64, error)
}

func New(userEditor userEditor) *UserSrv {
	return &UserSrv{userEditor: userEditor,
		usersMap: map[int64]struct{}{},
		mu:       sync.Mutex{}}
}

// CheckUser проверяет наличие пользователя
func (s *UserSrv) CheckUser(ctx context.Context, tgID int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.usersMap[tgID]
	logrus.Debugf("user %d found in map: %v", tgID, ok)
	return ok
}

// SaveUser сохраняет пользователя в мапе s.usersMap и в БД
func (s *UserSrv) SaveUser(ctx context.Context, tgID int64) error {
	logrus.Debugf("saving user %d", tgID)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.usersMap[tgID] = struct{}{}

	return s.userEditor.Save(ctx, tgID)
}

// LoadUsers загружает в память всех пользователей из БД
func (s *UserSrv) LoadUsers(ctx context.Context, users []int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range users {
		s.usersMap[u] = struct{}{}
		logrus.Debugf("saved user %d", u)
	}

	return nil
}

func (s *UserSrv) GetAllUsers(ctx context.Context) ([]int64, error) {
	return s.userEditor.GetAll(ctx)
}
