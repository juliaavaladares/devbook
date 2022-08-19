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

func (p Publicacoes) BuscaPublicacao(publicacaoId int64) (modelos.Publicacao, error) {
	var publicacao modelos.Publicacao

	scripts := IniciaScripts()
	query := scripts.BuscaPublicacao

	linha, err := p.db.Query(query, publicacaoId)
	if err != nil {
		return publicacao, err
	}
	defer linha.Close()

	if linha.Next() {
		if err = linha.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); err != nil {
			return publicacao, err
		}
	}

	return publicacao, nil
}
func (p Publicacoes) BuscaPublicacoes(publicacaoId int64) ([]modelos.Publicacao, error) {
	var publicacoes []modelos.Publicacao

	scripts := IniciaScripts()
	query := scripts.BuscaPublicacoes

	linhas, err := p.db.Query(query, publicacaoId, publicacaoId)
	if err != nil {
		return publicacoes, err
	}
	defer linhas.Close()

	for linhas.Next() {
		var publicacao modelos.Publicacao
		err := linhas.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		)
		if err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)

	}

	return publicacoes, nil
}
