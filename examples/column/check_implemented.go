package main

import "context"

// 型 *UserUseCase が UserRepositoryインタフェース を実装していることをコンパイル時に保証する
var _ UserRepository = (*UserUseCase)(nil)

type UserRepository interface {
	Create(ctx context.Context, name string) error
}

type UserUseCase struct {
	userRepo UserRepository
}

func (u *UserUseCase) Create(ctx context.Context, name string) error {
	// TODO: ユーザを作成する
	return nil
}

func main() {
}
