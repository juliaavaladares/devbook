package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       int64     `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("nome é obrigatório e nao pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("nick é obrigatório e nao pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("email é obrigatório e nao pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("senha é obrigatório e nao pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
