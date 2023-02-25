package usecases

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRepository struct {
	userRepository repositories.ICreateUserRepository
}

type UserClaims struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	ExpiresAt int64  `json:"exp"`
	jwt.StandardClaims
}

func (userRepo *CreateUserRepository) Login(email, password string) (string, error) {
	user := &entities.User{}

	_, err := userRepo.userRepository.FindByEmail(email)

	if err != nil {
		return "", fmt.Errorf("Email ou palavra passe errada...")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("Email ou palavra passe errada..., ")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &UserClaims{
		ID:        user.ID,
		Email:     email,
		ExpiresAt: expirationTime.Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %v", err)
	}

	return tokenString, nil
}
