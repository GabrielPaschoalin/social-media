package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func ValidarToken(r *http.Request) error {
	token := extrairToken(r)

	if _, erro := verificarToken(token); erro != nil {
		return erro
	}

	return nil
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
	// token retorna "Bearer djkljdaskljdlka"
}

func verificarToken(token string) (*jwt.Token, error) {

	tokenParsed, erro := jwt.Parse(token, retornarChaveVerificacao)

	if erro != nil {
		return nil, erro
	}

	// Extrair as claims do token
	_, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return nil, errors.New("Token inválido")
	}

	return tokenParsed, nil

}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		// Não usou o método que foi usado na função acima para gerar o token
		return nil, errors.New("Método de assinatura inexperado")
	}

	return config.SecretKey, nil
}

func ExtrairIdUsuario(r *http.Request) (uint64, error) {
	token := extrairToken(r)

	tokenParsed, erro := verificarToken(token)
	if erro != nil {
		return 0, erro
	}

	permissoes, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return 0, errors.New("Token inválido")
	}

	usuarioId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
	if erro != nil {
		return 0, erro
	}

	return usuarioId, nil

}
