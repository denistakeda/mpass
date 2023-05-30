package auth_service

import (
	"context"

	"github.com/denistakeda/mpass/internal/domain"
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type (
	authService struct {
		secret string

		logger    zerolog.Logger
		userStore userStore
	}

	userStore interface {
		AddNewUser(ctx context.Context, login, passwordHash string) error
		GetUser(ctx context.Context, login string) (domain.User, error)
	}
)

type NewAuthServiceParams struct {
	Secret string

	LogService ports.LogService
	UserStore  userStore
}

func New(params NewAuthServiceParams) *authService {
	return &authService{
		secret: params.Secret,

		logger:    params.LogService.ComponentLogger("authService"),
		userStore: params.UserStore,
	}
}

func (a *authService) SignUp(ctx context.Context, login, password string) (string, error) {
	if login == "" {
		return "", errors.New("login is empty")
	}

	if password == "" {
		return "", errors.New("password is empty")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate password hash")
	}

	if err := a.userStore.AddNewUser(ctx, login, string(passwordHash)); err != nil {
		a.logger.Error().Err(err).Msg("failed to add a new user")
		return "", errors.Errorf("login %q is busy", login)
	}

	return a.generateJWT(login)
}

func (a *authService) SignIn(ctx context.Context, login, password string) (string, error) {
	if login == "" {
		return "", errors.New("login is empty")
	}
	if password == "" {
		return "", errors.New("password is empty")
	}

	user, err := a.userStore.GetUser(ctx, login)
	if err != nil {
		return "", errors.Wrap(err, "login or password incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.Wrap(err, "login or password incorrect")
	}
	return a.generateJWT(login)
}

func (a *authService) AuthenticateUser(ctx context.Context, token string) (domain.User, error) {
	login, err := a.extractLoginFromToken(token)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "failed to authenticate user")
	}

	user, err := a.userStore.GetUser(ctx, login)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "no such user")
	}

	return user, nil
}

func (a *authService) generateJWT(login string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", errors.Wrap(err, "failed to create a token")
	}

	return tokenString, nil
}

func (a *authService) extractLoginFromToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.secret), nil
	})

	if err != nil {
		return "", errors.Wrap(err, "failed to parse token")
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims["login"].(string), nil
	} else {
		return "", errors.New("failed to parse token")
	}
}
