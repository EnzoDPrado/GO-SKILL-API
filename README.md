# 🚀 Go API Study Template

Bem-vindo! Este é o meu repositório de estudos e consolidado de conceitos em **Golang**. Ele funciona como um "laboratório vivo": conforme avanço no ecossistema Go, atualizo este template com as melhores práticas e padrões que aprendo.

## 🎯 Objetivo

O projeto visa fornecer uma estrutura base robusta para APIs REST, focando em **Clean Architecture**, escalabilidade e segurança.

## 🛠️ Tecnologias e Padrões Implementados

- **Framework Web:** [Gin Gonic](https://github.com/gin-gonic/gin) para roteamento de alta performance.
- **Banco de Dados:** [GORM](https://gorm.io/) para persistência e manipulação de dados com PostgreSQL.
- **Arquitetura:** Separação clara em camadas (Domain, Use Cases, Handlers, Infra).
- **Injeção de Dependência:** Uso de interfaces para desacoplamento de serviços e repositórios.
- **Segurança:** Criptografia de senhas com **Bcrypt**.
  - Autenticação via **JWT (JSON Web Tokens)**.
- **Middlewares:** Interceptores customizados para Autenticação e Autorização (RBAC).
- **Containerização:** Configuração completa com **Docker** e **Docker Compose**.

## 🛣️ Rotas da API (v1)

Atualmente a API possui as seguintes rotas implementadas:

| Método  | Rota                     | Descrição                               | Proteção    |
| :------ | :----------------------- | :-------------------------------------- | :---------- |
| `POST`  | `/api/v1/auth/login`     | Realiza login e retorna cookie/token    | Público     |
| `POST`  | `/api/v1/users`          | Cadastro de novos usuários              | Público     |
| `GET`   | `/api/v1/users`          | Listagem de todos os usuários           | Autenticado |
| `DELETE`| `/api/v1/users/:id`      | Soft delete em um usuário               | Admin Only  |
| `PATCH` | `/api/v1/users/:id/role` | Altera o nível de acesso de um usuário  | Admin Only  |
