package helper

import (
	"Kanbanboard/domain"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TokenSuite struct {
	suite.Suite
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, &TokenSuite{})
}

func (s *TokenSuite) SetupTest() {
	os.Setenv("JWT_SECRET_KEY", "RAHASIA")
}

func (s *TokenSuite) TestGenerateToken_Success() {
	id := int64(1)
	role := "member"
	token, err := GenerateToken(id, role)

	s.NotNil(token)
	s.Nil(err)
}

func (s *TokenSuite) TestVerifyToken_Success() {
	id := int64(1)
	role := "member"
	token, err := GenerateToken(id, role)
	s.Require().NotNil(token)
	s.Require().Nil(err)

	claims, err := VerifyToken(token)
	s.Nil(err)
	s.NotNil(claims)

	s.NotNil(claims["id"])
	s.NotNil(claims["role"])

	// Float64 karena default type dari claims["id"] adalah float64
	s.Equal(float64(id), claims["id"])
	s.Equal("member", claims["role"])
}

func (s *TokenSuite) TestVerifyToken_InvalidSigningMethod() {
	token := "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njg5NTExODgsImlhdCI6MTY2ODM0NjM4OCwiZW1haWwiOiJoOEBoOC5jb216eCIsImlkIjo2NX0.Pj7Kvw0hINJeZajKbNU-qGz7Hy7c6wcOjotevkVJmHHGeooQxBDxCFN9q2DNyATH"
	expError := domain.ErrUnauthorized

	claims, err := VerifyToken(token)

	s.Nil(claims)
	s.NotNil(err)

	s.Equal(expError.Error(), err.Error())
}
