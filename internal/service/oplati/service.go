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
	GetUser(ctx context.Context, userId uuid.UUID) (domain.UserInfo, error)
	GetAllUsers(ctx context.Context) ([]domain.UserInfo, error)
	Transfer(ctx context.Context, userIDFirst uuid.UUID, userIDSecond uuid.UUID, amount int) error
	ChangeBalance(ctx context.Context, userId uuid.UUID, amount int, fn func(balance int, amount int) (int, error)) (domain.UserInfo, error)
	Deposit(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error)
	Withdraw(ctx context.Context, userId uuid.UUID, amount int) (domain.UserInfo, error)

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
		return []domain.UserInfo{}, err
	}

	return users, nil
}

func (s *Service) Transfer(ctx context.Context, userIDFirst uuid.UUID, userIDSecond uuid.UUID, amount int) error {
	err := s.db.Transfer(ctx, userIDFirst, userIDSecond, amount)

	if err != nil {
		return err
	}
	return nil
}
