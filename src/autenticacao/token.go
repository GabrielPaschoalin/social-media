package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}

	// Define lista de parâmetros para definir o token
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // Expira em 6 horas
	permissoes["usuarioId"] = usuarioID

	// Criar um token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	// Assinar o token
	return token.SignedString(config.SecretKey)
}
