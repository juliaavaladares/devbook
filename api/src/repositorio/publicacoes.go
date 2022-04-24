package repositorio

import (
	"database/sql"
	"devbook-api/src/modelos"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db: db}
}

func (p Publicacoes) CriaPublicacao(publicacao modelos.Publicacao) (int64, error) {
	scripts := IniciaScripts()
	query := scripts.CriaPublicacao

	statement, err := p.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if err != nil {
		return 0, err
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ultimoIdInserido, nil
}
