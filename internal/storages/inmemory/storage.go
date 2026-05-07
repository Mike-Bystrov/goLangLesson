package inmemory

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/thnxvlad/oplati/internal/domain"
)

type Storage struct {
	db map[uuid.UUID]domain.UserInfo
	sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		db: make(map[uuid.UUID]domain.UserInfo),
	}
}

func (s *Storage) CreateUser(ctx context.Context, ui domain.UserInfo) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.db[ui.Id]; ok {
		return errors.New("id already exists")
	}

	s.db[ui.Id] = ui

	return nil
}

func (s *Storage) Deposit(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error) {
	s.Lock()
	defer s.Unlock()

	ui, ok := s.db[userId]
	if !ok {
		return domain.UserInfo{}, errors.New("user id does not exist")
	}

	ui.Balance += amount

	s.db[userId] = ui

	return ui, nil
}

func (s *Storage) Withdraw(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error) {
	s.Lock()
	defer s.Unlock()

	ui, ok := s.db[userId]
	if !ok {
		return domain.UserInfo{}, errors.New("user id does not exist")
	}

	ui.Balance -= amount

	s.db[userId] = ui

	return ui, nil
}

func (s *Storage) GetUser(ctx context.Context, userId uuid.UUID) (domain.UserInfo, error) {
	s.RLock()
	defer s.RUnlock()

	ui, ok := s.db[userId]
	if !ok {
		return domain.UserInfo{}, errors.New("user id does not exist")
	}

	return ui, nil
}

func (s *Storage) GetAllUsers(ctx context.Context) ([]domain.UserInfo, error) {
	s.RLock()
	defer s.RUnlock()

	users := make([]domain.UserInfo, 0, len(s.db))

	for _, ui := range s.db {
		users = append(users, ui)
	}

	return users, nil
}

func (s *Storage) Transfer(ctx context.Context, userIDFrom uuid.UUID, userIDTo uuid.UUID, amount int) ([]domain.UserInfo, error) {
	s.Lock()
	defer s.Unlock()
	ui1, ok1 := s.db[userIDFrom]
	ui2, ok2 := s.db[userIDTo]

	if !ok1 {
		return []domain.UserInfo{}, errors.New("there is no user with id: " + userIDFrom.String())
	}
	if !ok2 {
		return []domain.UserInfo{}, errors.New("there is no user with id: " + userIDTo.String())
	}

	if ui1.Balance-amount < 0 {
		return []domain.UserInfo{}, errors.New("Not enough money for transfer")
	}

	s.Withdraw(ctx, userIDFrom, amount)
	s.Deposit(ctx, userIDTo, amount)

	return []domain.UserInfo{ui1, ui2}, nil
}
