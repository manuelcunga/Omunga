package usecases

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCases struct {
	userRepository repositories.ICreateUserRepository
}

type UserClaims struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	ExpiresAt int64  `json:"exp"`
	jwt.StandardClaims
}

func NewLoginUseCases(userRepo repositories.ICreateUserRepository) LoginUseCases {
	return LoginUseCases{userRepository: userRepo}
}

func (loginUseCases *LoginUseCases) Login(data *dtos.LoginDTO) (string, error) {
	user, err := loginUseCases.userRepository.FindByEmail(data.Email)

	if err != nil {
		return "", fmt.Errorf("Email ou palavra passe errada...")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return "", fmt.Errorf("Email ou palavra passe errada..., ")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &UserClaims{
		ID:        user.ID,
		Email:     data.Email,
		ExpiresAt: expirationTime.Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %v", err)
	}

	return tokenString, nil
}
