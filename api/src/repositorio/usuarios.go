package repositorio

import (
	"database/sql"
	"devbook-api/src/modelos"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db: db}
}

func (u Usuarios) CriaUsuario(usuario modelos.Usuario) (int64, error) {
	scripts := IniciaScripts()
	query := scripts.CriaUsuario
	statement, err := u.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ultimoIdInserido, nil
}

func (u Usuarios) BuscaUsuario(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	scripts := IniciaScripts()
	query := scripts.BuscaUsuarios

	linhas, err := u.db.Query(query, nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm)
		if err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
