package usecase

import (
	"errors"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/user"
	"fmt"
)

// 入力DTO
type CreateUserInput struct {
	Name  string
	Email string
}

// 出力DTO
type CreateUserOutput struct {
	Id                 string
	Name               string
	Email              string
	Status             string
	CurrentBorrowCount int
	OverdueFees        float64
}

type CreateUserUseCase struct {
	// コンストラクター: リポジトリインターフェースに依存
	userRepository user.UserRepository
}

func NewCreateUserUseCase(repo user.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: repo}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInput) (*CreateUserOutput, error) {
	// ビジネスルール: 重複メールをチェック
	existingUser, err := uc.userRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check for existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// ファクトリーメソッドでユーザーを作成
	u := user.NewUser(input.Name, input.Email)

	// リポジトリに永続化
	if err := uc.userRepository.Save(u); err != nil {
		return nil, err
	}

	// DTOを返す
	return &CreateUserOutput{
		Id:                 u.Id().Value(),
		Name:               u.Name(),
		Email:              u.Email(),
		Status:             string(u.Status()),
		CurrentBorrowCount: u.CurrentBorrowCount(),
		OverdueFees:        u.OverdueFees(),
	}, nil
}