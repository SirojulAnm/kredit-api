package auth

import (
	"errors"
	"kredit-api/log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(tokenStr string) (*CustomClaims, error)
}

type jwtService struct {
	SecretKey string
}

type CustomClaims struct {
	jwt.StandardClaims
}

var SECRET_KEY = []byte("s3c.r3t_s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	now := time.Now().UTC()
	//24 * time.Hour * 30
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 0,
			IssuedAt:  now.Unix(),
			Subject:   strconv.Itoa(userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		log.Error(&logrus.Fields{"error": err.Error()}, "Error while encode jwt")
		return "", err
	}

	return tokenStr, nil
}

func (s *jwtService) ValidateToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		log.Error(&logrus.Fields{"error": err.Error()}, "Error while parse jwt")
		return nil, errors.New("Unauthorized, Error while parse jwt")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Error(nil, "Error while claims jwt")
		return nil, errors.New("Unauthorized, Error while claims jwt")
	}
}
