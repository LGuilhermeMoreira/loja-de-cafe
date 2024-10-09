# Backend API em Go

Este projeto é uma API backend desenvolvida em Go usando as bibliotecas Gin Gonic e GORM. Ele implementa autenticação JWT e utiliza criptografia de senha para login seguro.

## Funcionalidades

- **Autenticação JWT**: Geração e validação de tokens JWT para autenticação segura.
- **Criptografia de senha**: As senhas dos usuários são criptografadas para garantir a segurança dos dados.
- **CRUD de Usuários**: Implementação básica de criação, leitura, atualização e exclusão de usuários usando GORM.
- **Rotas Protegidas**: Algumas rotas exigem autenticação JWT para acesso.
- **Banco de dados relacional**: Integração com um banco de dados SQL usando o ORM GORM.
## Tecnologias

- **Linguagem**: Go
- **Framework**: Gin Gonic
- **ORM**: GORM
- **Autenticação**: JWT (JSON Web Tokens)
- **Criptografia de Senha**: bcrypt
