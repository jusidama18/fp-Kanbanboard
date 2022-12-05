package usecase

import (
	"context"
	"fmt"
	"time"

	"Kanbanboard/app/helper"
	"Kanbanboard/domain"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepository: u}
}

func (u *UserUsecase) Login(ctx context.Context, user *domain.User) (token string, err error) {
	res, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		err = domain.ErrUnauthorized
		return
	}
	comparePass := helper.ComparePass([]byte(res.Password), []byte(user.Password))
	if !comparePass {
		err = domain.ErrUnauthorized
		return
	}
	token, err = helper.GenerateToken(res.ID, res.Role)
	if err != nil {
		return
	}
	return
}

func (u *UserUsecase) Register(ctx context.Context, user *domain.User) (domain.User, error) {
	_, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return domain.User{}, domain.ErrConflict
	}

	user.Password, err = helper.HashPass(user.Password)
	user.Role = "member"
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}

	userId, err := u.userRepository.StoreUser(ctx, user)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	userData, err := u.userRepository.GetUserByID(ctx, userId)
	if err != nil {
		return domain.User{}, domain.ErrNotFound
	}
	return userData, nil
}

func (u *UserUsecase) RegisterAdmin(ctx context.Context, user *domain.User) (domain.User, error) {
	fmt.Println("admin")
	_, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return domain.User{}, domain.ErrConflict
	}

	user.Password, err = helper.HashPass(user.Password)
	user.Role = "admin"
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}

	userId, err := u.userRepository.StoreUser(ctx, user)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	userData, err := u.userRepository.GetUserByID(ctx, userId)
	if err != nil {
		return domain.User{}, domain.ErrNotFound
	}
	return userData, nil
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) (domain.User, error) {
	_, err := u.userRepository.GetUserByID(ctx, user.ID)
	if err != nil {
		return domain.User{}, domain.ErrNotFound
	}
	user.UpdatedAt = time.Now()
	err = u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	return *user, nil
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	_, err := u.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.ErrNotFound
	}
	err = u.userRepository.DeleteUser(ctx, id)
	return err
}
