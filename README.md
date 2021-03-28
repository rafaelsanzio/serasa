<h1 align="center">
  <img style="background-color: #312e38; border-radius: 10px;" alt="serasa-logo" src="https://www.foregon.com/blog/wp-content/uploads/2020/10/servicosdaserasaexperian-1068x601.jpeg" />
</h1>

## 🔖 Sobre o projeto

O projeto **Serasa Negativations**. Tendo como objetivo, criar uma API para importação de negativações e fazer consultas com filtros.

## 💻 Tecnologias

- <img width="20px" src="https://img.icons8.com/color/2x/golang.png" /> [Golang](https://golang.org/ "Golang")
- <img width="20px" src="https://img.icons8.com/dusk/2x/docker.png" /> [Docker](https://www.docker.com/ "Docker")

## ▶️ Getting Started

- **Passo 1️⃣** : git clone do projeto [Serasa](https://github.com/rafaelsanzio/serasa "Serasa")
- **Passo 2️⃣** : executar a instalação do [Go](https://golang.org/ "Go") e [Docker](https://www.docker.com/ "Docker")

```bash
   # Navegando até a pasta do projeto
   $ cd serasa

   # Instalando as dependências do projeto
   $ go get .

   # Criando imagem do banco de dados (mongoDB) da aplicação
   $ docker run --name serasa -p 27017:27017 -d -t mongo

   # Iniciando banco de dados da aplicação
   $ docker start serasa

   # Criar arquivo .env baseado no .env.example
   # Preencher o arquivo .env
   # Gerar token de acordo com o API_KEY

   # Rodando a aplicação
   $ go run main.go
```

## ⚙️ Exemplificando rotas

- [Documentação da API](https://documenter.getpostman.com/view/15147727/TzCJgpwK "Documentação da API")

## ㊗ ️ Considerações

- Projeto desenvolvido by:

  - <a href="https://github.com/rafaelsanzio">
      <img src="https://img.shields.io/badge/-Rafael%20Sanzio-000000?style=flat&logo=GitHub&logoColor=#000000" />
    </a>

  - <a href="https://www.linkedin.com/in/rafael-sanzio-012778143/">
      <img src="https://img.shields.io/badge/-Rafael%20Sanzio-0077B5?style=flat&logo=LinkedIN&logoColor=#000000" />
    </a>
