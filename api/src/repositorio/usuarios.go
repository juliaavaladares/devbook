package repositorio

import (
	"database/sql"
	"devbook-api/src/modelos"
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
