package repositorio

type Scripts struct {
	CriaUsuario          string
	BuscaUsuarios        string
	BuscaUsuarioPorId    string
	AtualizaUsuario      string
	DeletaUsuario        string
	BuscaUsuarioPorEmail string
	SeguirUsuario        string
	PararDeSeguirUsuario string
	BuscaSeguidores      string
	BuscaSeguindo        string
	BuscaSenha           string
	AtualizaSenha        string
}

func IniciaScripts() Scripts {
	scripts := Scripts{}

	scripts.CriaUsuario = `insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)`
	scripts.BuscaUsuarios = `select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?`
	scripts.BuscaUsuarioPorId = `select id, nome, nick, email, criadoEm from usuarios where id = ?`
	scripts.AtualizaUsuario = `update usuarios set nome = ?, nick = ?, email = ?, where id = ?`
	scripts.DeletaUsuario = `delete from usuarios where id = ?`
	scripts.BuscaUsuarioPorEmail = `select id, senha from usuarios where email = ?`
	scripts.SeguirUsuario = `insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)`
	scripts.PararDeSeguirUsuario = `delete from seguidores where usuario_id = ? and seguidor_id = ?`
	scripts.BuscaSeguidores = ` select u.id, u.nome,
									u.nick, u.email, u.criadoEm
  								from
									usuarios u
									inner join seguidores on u.id = s.seguidor_id
  								where
									usuario_id = ?`
	scripts.BuscaSeguindo = ` select u.id, u.nome,
									u.nick, u.email, u.criadoEm
  								from
									usuarios u
									inner join seguidores on u.id = s.usuario_id
  								where
									usuario_id = ?`
	scripts.BuscaSenha = `select senha from usuarios where id = ?`
	scripts.BuscaSenha = `update usuarios set senha = ? where id = ?`

	return scripts
}
