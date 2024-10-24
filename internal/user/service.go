package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	. "github.com/mlgnu/cocode/internal/user/repository"
	authmiddleware "github.com/mlgnu/cocode/pkg"
)

type Service struct {
	Queries *Queries
}

func NewService(queries *Queries) *Service {
	return &Service{Queries: queries}
}

func (s *Service) GetUser(ctx context.Context, id int32) (User, error) {
	return s.Queries.GetUser(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, req UpdateUserRequest) error {
	avatar := pgtype.Text{
		String: req.Avatar,
		Valid:  req.Avatar != "",
	}

	err := s.Queries.UpdateUser(ctx, UpdateUserParams{
		ID:        req.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Avatar:    avatar,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id int32) error {
	err := s.Queries.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ChangePassword(ctx context.Context, id int32) error {
	user := ctx.Value("user").(*authmiddleware.User)
	fmt.Println(user)
	return nil
	// s.Queries.GetUser(, 1)
	// err := s.Queries.ChangePassword(context.Context, arg ..ChangePasswordParams)
}
