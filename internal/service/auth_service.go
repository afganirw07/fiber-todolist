package service

import (
	"errors"
	"time"
	"todolist-backend/internal/config"
	"todolist-backend/internal/model"
	"todolist-backend/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}
type UserService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *AuthService) Register(user *model.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)
	return s.Repo.Create(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	var user model.User
	if err := s.Repo.FindByUsername(username, &user); err != nil {
		return "", errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid password or username")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.FindAll()
}
