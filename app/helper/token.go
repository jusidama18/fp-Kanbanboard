package helper

import (
	"os"

	"Kanbanboard/domain"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(id int64, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))

	return signedToken, err
}

func VerifyToken(headerToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(headerToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnauthorized
		} else if method != jwt.SigningMethodHS256 {
			return nil, domain.ErrUnauthorized
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, domain.ErrUnauthorized
	}

	return claims, nil
}
