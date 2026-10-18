package seguranca

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {

	// Recebe uma string e coloca um hash nela
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	return bytes, err
}

func VerificarSenha(senhaComHash, senhaString string) error {
	// Compara uma senha com o hash e retorna se elas são iguais
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
