package helper

import (
	"os"
	"strings"

	"Kanbanboard/domain"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(id int64, role string) string {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(headerToken string) (jwt.MapClaims, error) {
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, domain.ErrUnauthorized
	}
	if len(headerToken) <= 6 {
		return nil, domain.ErrUnauthorized
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnauthorized
		}
		return []byte(secretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, domain.ErrUnauthorized
	}

	return claims, nil
}
