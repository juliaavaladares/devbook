package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	Id        int64     `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   int64     `json:"autorId,omitempty"`
	AutorNick int64     `json:"autorNick,omitempty"`
	Curtidas  int64     `json:"curtidas,omitempty"`
	CriadoEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.validar(); err != nil {
		return err
	}

	err := publicacao.formatar()
	if err != nil {
		return err
	}

	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("titulo é obrigatório e nao pode estar em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("conteudo é obrigatório e nao pode estar em branco")
	}

	return nil
}

func (publicacao *Publicacao) formatar() error {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)

	return nil
}
