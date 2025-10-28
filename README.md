
# Teste Automatizado de API com Golang

Este projeto tem como objetivo ajudar iniciantes a entender como criar testes automatizados de API REST utilizando a linguagem Go (Golang).



## 📚 Sobre

O repositório contém um exemplo simples de teste automatizado para a API pública [FakeRestAPI](https://fakerestapi.azurewebsites.net/api/v1/Activities) Activities, demonstrando como:

- Fazer requisições HTTP em Go
- Validar status code, headers e conteúdo da resposta
- Validar o contrato dos dados retornados

## 🛠️ Bibliotecas Utilizadas

- [resty](https://github.com/go-resty/resty) — Cliente HTTP para Go
- [testify/assert](https://github.com/go-resty/resty) — Biblioteca de asserções para testes
- encoding/json — Biblioteca padrão do Go para manipulação de JSON
- time — Biblioteca padrão do Go para manipulação de datas


## 🌐 API Testada

Endpoint: https://fakerestapi.azurewebsites.net/api/v1/Activities

Método: GET

## 🚀 Como executar o teste

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/seu-repo.git
cd seu-repo
```

2. Instale as dependências:

```bash
go get github.com/go-resty/resty/v2
go get github.com/stretchr/testify/assert
```

3. Execute os testes:

```bash
go test -v
```

O teste irá realizar uma requisição GET para a API, validar o status da resposta, o formato dos dados e se os campos obrigatórios estão presentes e corretos.

