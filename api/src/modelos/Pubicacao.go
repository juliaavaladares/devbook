package modelos

import "time"

type Publicacao struct {
	Id        int64     `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   int64     `json:"autorId,omitempty"`
	AutorNick int64     `json:"autorNick,omitempty"`
	Curtidas  int64     `json:"curtidas,omitempty"`
	CriadoEm  time.Time `json:"criadaEm,omitempty"`
}
