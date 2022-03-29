package repositorio

type Scripts struct {
	CriaUsuario       string
	BuscaUsuarios     string
	BuscaUsuarioPorId string
	AtualizaUsuario   string
}

func IniciaScripts() Scripts {
	scripts := Scripts{}

	scripts.CriaUsuario = `insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)`
	scripts.BuscaUsuarios = `select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?`
	scripts.BuscaUsuarioPorId = `select id, nome, nick, email, criadoEm from usuarios where id = ?`
	scripts.AtualizaUsuario = `update usuarios set nome = ?, nick = ?, email = ?, where id = ?`

	return scripts
}
