package oplati

import (
	"context"

	"github.com/google/uuid"
	"github.com/thnxvlad/oplati/internal/domain"
)

type Service struct {
	db OplatiDatabase
}

func New(db OplatiDatabase) *Service {
	return &Service{
		db: db,
	}
}

type OplatiDatabase interface {
	CreateUser(ctx context.Context, ui domain.UserInfo) error
	Deposit(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error)
	GetUser(ctx context.Context, userId uuid.UUID) (domain.UserInfo, error)
	Withdraw(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error)
	GetAllUsers(ctx context.Context) ([]domain.UserInfo, error)
	Transfer(ctx context.Context, userIDFirst uuid.UUID, userIDSecond uuid.UUID, amount int) ([]domain.UserInfo, error)
}

func (s *Service) CreateUser(ctx context.Context, name string) (domain.UserInfo, error) {
	ui := domain.UserInfo{
		Id:      uuid.New(),
		Name:    name,
		Balance: 0,
	}

	err := s.db.CreateUser(ctx, ui)
	if err != nil {
		return domain.UserInfo{}, err
	}

	return ui, nil
}

func (s *Service) Deposit(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error) {
	ui, err := s.db.Deposit(ctx, userId, amount)
	if err != nil {
		return domain.UserInfo{}, err
	}

	return ui, nil
}

func (s *Service) Withdraw(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error) {
	ui, err := s.db.Withdraw(ctx, userId, amount)
	if err != nil {
		return domain.UserInfo{}, err
	}

	return ui, nil
}

func (s *Service) GetUser(ctx context.Context, userId uuid.UUID) (domain.UserInfo, error) {
	ui, err := s.db.GetUser(ctx, userId)
	if err != nil {
		return domain.UserInfo{}, err
	}

	return ui, nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]domain.UserInfo, error) {
	users, err := s.db.GetAllUsers(ctx)

	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *Service) Transfer(ctx context.Context, userIDFirst uuid.UUID, userIDSecond uuid.UUID, amount int) ([]domain.UserInfo, error) {
	userFirst, err1 := s.db.GetUser(ctx, userIDFirst)
	userSecond, err2 := s.db.GetUser(ctx, userIDSecond)

	if err1 != nil || err2 != nil {
		return []domain.UserInfo{userFirst, userSecond}, err1
	}

	return []domain.UserInfo{userFirst, userSecond}, nil
}
