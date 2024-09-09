package auth

import (
	"errors"
	"log"

	"github.com/dev-oleksandrv/db"
	jwtutils "github.com/dev-oleksandrv/internal/jwt-utils"
	"github.com/dev-oleksandrv/repository"
	"golang.org/x/crypto/bcrypt"
)  

type AuthService struct {
	userRepo repository.GORMUserRepository
}

func NewAuthService(userRepo repository.GORMUserRepository) *AuthService {
	return &AuthService{ userRepo }
}

func (s *AuthService) RegisterUser(email, password string) (*db.User, error) {
	if user := s.userRepo.GetUserByEmail(email); user.ID != 0 {
		log.Printf("User with such email found: %v", email)
		return nil, errors.New("user with such email is found")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}
	user, err := s.userRepo.Create(db.User{ Email: email, Password: string(hashedPassword) })
	if err != nil {
		log.Printf("Error while creating a user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) LoginUser(email, password string) (string, string, error) {
	user := s.userRepo.GetUserByEmail(email)
	if user.ID == 0 {
		log.Printf("User with such email not found: %v", email)
		return "", "", errors.New("user with such email is found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		log.Println("Password is invalid")
		return "", "", errors.New("user password is invalid")
	}
	accessToken, refreshToken, err := jwtutils.GenerateTokens(user.ID, user.Email)
	if err != nil {
		log.Println("Cannot generate a JWT token")
		return "", "", err
	}
	log.Printf("tokens: %v %v", accessToken, refreshToken)
	return accessToken, refreshToken, nil
}