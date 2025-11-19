package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
	"gitlab.com/genuicomm/api/domain"
	"gitlab.com/genuicomm/api/modules/user"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo user.UserRepository
}

// NewAuthUsecase membuat instance baru dari AuthUsecase
func NewAuthUsecase(userRepo user.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

// Login melakukan autentikasi pengguna dan menghasilkan token
func (u *authUsecase) Login(email, password string) (domain.LoginResponse, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return domain.LoginResponse{}, errors.New("invalid credentials")
	}

	accessToken, err := u.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		User:        user,
	}, nil
}

func (u *authUsecase) GenerateToken(userID string, email string) (string, error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-256-bit-secret" // Default secret, should be changed in production
	}

	return token.SignedString([]byte(jwtSecret))
}

func (u *authUsecase) GenerateRefreshToken(userID string, email string) (string, error) {
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"type":    "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-256-bit-secret"
	}

	return token.SignedString([]byte(jwtSecret))
}
