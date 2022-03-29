package repositorio

type Scripts struct {
	CriaUsuario          string
	BuscaUsuarios        string
	BuscaUsuarioPorId    string
	AtualizaUsuario      string
	DeletaUsuario        string
	BuscaUsuarioPorEmail string
}

func IniciaScripts() Scripts {
	scripts := Scripts{}

	scripts.CriaUsuario = `insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)`
	scripts.BuscaUsuarios = `select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?`
	scripts.BuscaUsuarioPorId = `select id, nome, nick, email, criadoEm from usuarios where id = ?`
	scripts.AtualizaUsuario = `update usuarios set nome = ?, nick = ?, email = ?, where id = ?`
	scripts.DeletaUsuario = `delete from usuarios where id = ?`
	scripts.BuscaUsuarios = `select id, senha from usuarios where email = ?`

	return scripts
}
