package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	. "github.com/mlgnu/cocode/internal/auth/repository"
	authmiddleware "github.com/mlgnu/cocode/pkg"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Queries *Queries
}

func NewService(queries *Queries) *Service {
	return &Service{Queries: queries}
}

func (s *Service) Register(ctx context.Context, req RegisterRequest) error {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	s.Queries.AddUser(ctx, AddUserParams{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})

	return nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func signToken(user GetUserAuthRow) (string, error) {
	// claims := jwt.MapClaims{
	// 	"Id":    user.ID,
	// 	"email": user.Email,
	// 	"role":  user.Role,
	// 	"exp":   time.Now().Add(time.Hour * 24).Unix(),
	// }
	claims := &authmiddleware.Claims{
		ID:    int(user.ID),
		Email: user.Email,
		Role:  string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	key := []byte("adfadsfasdfasdfasdaf")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	user, err := s.Queries.GetUserAuth(ctx, req.Email)
	if err != nil {
		return "", err
	}

	if !checkPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := signToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	return s.Queries.GetUserByEmail(ctx, email)
}

func (s *Service) GetUserById(ctx context.Context, id int32) (GetUserByIdRow, error) {
	return s.Queries.GetUserById(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, req UpdateUserRequest) error {
	s.Queries.UpdateUser(ctx, UpdateUserParams{
		ID:        req.Id,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	return nil
}
