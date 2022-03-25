package repositorio

type Scripts struct {
	CriaUsuario string
}

func IniciaScripts() Scripts {
	scripts := Scripts{}

	scripts.CriaUsuario = `insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)`

	return scripts
}
