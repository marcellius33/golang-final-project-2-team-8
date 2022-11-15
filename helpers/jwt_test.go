package helpers

import (
	"mygram/models"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

var (
	user        = &models.User{
		ID	    : 1,
		Username: "testing",
		Email   : "testing@gmail.com",
	}
)

func TestSuccessGenerateToken(t *testing.T) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	t.Logf("Berhasil generate token : %s", signedToken)
}

func TestFailedGenerateToken(t *testing.T) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString("")

	t.Logf("Gagal generate token : %s", signedToken)
}