package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"errors"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (rep Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := rep.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (rep Usuarios) BuscarUsuarios(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	// Buscar todos os usuários que atendem o nome ou nick
	linhas, erro := rep.db.Query(
		"SELECT id, nome, nick, email, criado_em FROM usuarios WHERE nome LIKE ? or nick LIKE ?",
		nomeOuNick,
		nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var u modelos.Usuario

		if erro = linhas.Scan(&u.ID, &u.Nome, &u.Nick, &u.Email, &u.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, u)
	}

	// Verifica se a iteração foi interrompida por algum erro
	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return usuarios, nil

}

func (rep Usuarios) BuscarUsuario(usuarioId uint64) (modelos.Usuario, error) {

	// Faz a consulta ao banco
	linhas, erro := rep.db.Query("SELECT id, nome, nick, email, criado_em from usuarios where id = ?", usuarioId)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	// Nenhuma linha: ou deu erro na leitura, ou o usuário não existe
	if !linhas.Next() {
		if erro = linhas.Err(); erro != nil {
			return modelos.Usuario{}, erro
		}
		return modelos.Usuario{}, errors.New("Usuário não encontrado")
	}

	// Salva a linha no usuário
	var u modelos.Usuario
	if erro = linhas.Scan(&u.ID, &u.Nome, &u.Nick, &u.Email, &u.CriadoEm); erro != nil {
		return modelos.Usuario{}, erro
	}

	return u, nil

}

func (rep Usuarios) AtualizarUsuario(usuarioId uint64, u modelos.Usuario) error {

	// Preparar a query
	statement, erro := rep.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	// Rodar a query
	_, erro = statement.Exec(u.Nome, u.Nick, u.Email, usuarioId)
	if erro != nil {
		return erro
	}

	return nil
}

func (rep Usuarios) DeletarUsuario(usuarioId uint64) error {
	// Preparar a query
	statement, erro := rep.db.Prepare("DELETE FROM usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	// Rodar a query
	_, erro = statement.Exec(usuarioId)
	if erro != nil {
		return erro
	}

	return nil

}

func (rep Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, erro := rep.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var u modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(&u.ID, &u.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return u, nil

}
