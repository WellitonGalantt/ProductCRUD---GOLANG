package auth

// go mod tidy para instalar os imports
import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Teste, depois colocar em env para mais seguranca
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,                                // subject = id do usuário
		"exp": time.Now().Add(24 * time.Hour).Unix(), // expira em 24h
		"iat": time.Now().Unix(),                     // emitido em
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, err
	}

	return token, claims, nil

}
