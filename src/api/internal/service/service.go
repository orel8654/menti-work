package service_api

import (
	"context"
	"errors"
	"fmt"
	"menti/pkg/types"
	b64 "encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secretLine = "middle"

type Repo interface {
	CreateUser(ctx context.Context, data types.UserPayloadCreat) error
	GetUser(ctx context.Context, username string) (user types.User, err error)
	GetBasicToken(ctx context.Context, IDkey int) (res types.UserToken, err error)
	CreateBasicToken(ctx context.Context, token string, IDkey int) (res types.UserToken, err error)
	GetUserByBasicToken(ctx context.Context, token string) (user types.User, err error)
}

type Service struct {
	repo       Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterUserService(ctx context.Context, credentials types.UserPayload) error {
	if err := s.repo.CreateUser(
		ctx,
		types.UserPayloadCreat{
			Username: credentials.Username,
			UUID:     PasswordEncoding(credentials.Username),
			Password: PasswordEncoding(credentials.Password),
		},
	); err != nil {
		return err
	}
	return nil
}

func (s *Service) LoginUserService(ctx context.Context, credentials types.UserPayload) (res types.User, err error) {
	res, err = s.repo.GetUser(ctx, credentials.Username)
	if err != nil {
		return res, err
	}
	if PasswordEncoding(credentials.Password) != res.Password {
		return res, errors.New("invalid password")
	}
	return res, nil
}

func (s *Service) LoginUserByBasicService(ctx context.Context, credentials types.UserPayload) (res types.UserToken, err error) {
	user, err := s.LoginUserService(ctx, credentials)
	if err != nil {
		return res, err
	}
	res, err = s.repo.GetBasicToken(ctx, user.IDKey)
	if err == nil {
		return res, err
	}
	token := TokenEncoding(credentials.Username, credentials.Password)
	res, err = s.repo.CreateBasicToken(ctx, token, user.IDKey)
	if err != nil {
		return res, err
	}
	return res, err
}

func (s *Service) LoginUserByBearerService(ctx context.Context, credentials types.UserPayload) (res types.UserToken, err error) {
	expareDate := time.Now().Add(10 * time.Minute).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = credentials.Username
	claims["password"] = credentials.Password
	claims["exp"] = expareDate
	tokenString, err := token.SignedString([]byte(secretLine))
	if err != nil {
		return res, err
	}
	return types.UserToken{Token: tokenString}, err
}

func (s *Service) AuthMiddlewareService(ctx context.Context, parts []string) (user types.User, err error) {
	if parts[0] == "Basic" {
		user, err = s.repo.GetUserByBasicToken(ctx, parts[1])
		if err != nil {
			return user, errors.New("can not found user")
		}
	} else if parts[0] == "Bearer" {
		username, err := VerifyJWTToken(parts[1])
		if err != nil {
			return user, err
		}
		return s.repo.GetUser(ctx, username)
	}
	return user, err
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func VerifyJWTToken(data string) (username string, err error) {
	token, err := jwt.ParseWithClaims(data, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretLine), nil
	})
	if err != nil {
		return username, errors.New("can not parse token")
	}
	if !token.Valid {
		return username, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return username, errors.New("token has expired")
	}
	return claims.Username, err
}

func TokenEncoding(username string, password string) string {
	token := fmt.Sprintf("%s:%s", username, password)
	return b64.StdEncoding.EncodeToString([]byte(token))
}

func PasswordEncoding(data string) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}
