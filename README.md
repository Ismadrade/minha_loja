
<div style="display: inline_block"><br>
  <img align="center" alt="Spring" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" />
  <img align="center" alt="Spring" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg" />
  <img align="center" alt="RabbitMQ" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/html5/html5-plain.svg" />
  <img align="center" alt="Spring" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/bootstrap/bootstrap-plain.svg" />   
</div>

# Minha Loja - MVC com Go Lang
Pequeno projeto desenvolvido em Go Lang, que permite criar, consultar, editar e deletar um produto.

## Requisitos

- Go Lang
- Postgres

## Configurando o Ambiente

### Postgres
1) Instalar o Postgres. [Download](https://www.postgresql.org/download/) e  [Instalação](https://www.postgresqltutorial.com/postgresql-getting-started/install-postgresql/)

2) Criar uma database chamada minha_loja


3) Criar uma tabela chamada produtos:
```postgresql
CREATE TABLE produtos (
	id SERIAL PRIMARY KEY,
	nome VARCHAR,
	descricao VARCHAR,
	preco DECIMAL,
	quantidade INTEGER
)
```

4) Inserir alguns dados de exemplo (Opcional):
```postgresql
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES 
('Casaco', 'Casaco de couro preto', 119.99, 5),
('Fone JBL', 'Fone com fio JBL', 89.99, 15);
```

### Go Lang

1) [Baixar](https://go.dev/doc/install) e [instalar](https://medium.com/@rafaelmoraisdev/como-instalar-go-no-windows-10-7787faac3a7f) o Go Lang.

## Executando o projeto
Dentro da pasta go_web_mvc, rodar o seguinte comando:
```shell
go run main.go  
```
O projeto estará rodando em localhost na porta 8000
```http request
http://localhost:8000/
```

## Referências
- [Alura](https://cursos.alura.com.br/course/go-lang-web)
