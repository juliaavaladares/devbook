CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;


CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores (
    usuario_id int not null,
    FOREIGN key (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN key (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null unique,

    autor_id int not null,  
    FOREIGN key (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

insert into publicacoes (titulo, conteudo, autor_id)
values
("Publicacao do usuario 1", "Publicacao usuario 1: Oba!", 1),
("Publicacao do usuario 2", "Publicacao usuario 2: Oba!", 2),
("Publicacao do usuario 3", "Publicacao usuario 3: Oba!", 3),
