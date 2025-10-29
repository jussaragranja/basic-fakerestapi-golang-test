
# Teste Automatizado de API com Golang

Este projeto tem como objetivo ajudar iniciantes a entender como criar testes automatizados de API REST utilizando a linguagem **Go (Golang)**.

---

## 📚 Sobre

O repositório contém um exemplo simples de teste automatizado para a API pública [FakeRestAPI](https://fakerestapi.azurewebsites.net/api/v1/Activities), demonstrando como:

- Fazer requisições HTTP em Go
- Validar status code, headers e conteúdo da resposta
- Validar o contrato dos dados retornados
- Garantir que os campos de data estão no formato **RFC3339**

O código está todo comentado para facilitar o entendimento e aprendizado.

---

## 🛠️ Bibliotecas Utilizadas

- [Resty](https://github.com/go-resty/resty) — Cliente HTTP para Go
- [Testify/assert](https://github.com/stretchr/testify) — Biblioteca de asserções para testes
- `encoding/json` — Biblioteca padrão do Go para manipulação de JSON
- `time` — Biblioteca padrão do Go para manipulação de datas

---

## 🌐 API Testada

- **Endpoint:** `https://fakerestapi.azurewebsites.net/api/v1/Activities`
- **Método:** `GET`

---

## 🚀 Como executar o teste

1. Clone o repositório:

```bash
git clone https://github.com/jussaragranja/basic-fakerestapi-golang-test.git
cd basic-fakerestapi-golang-test
```

2. Instale as dependências:

```bash
go mod tidy
```

3. Execute os testes:

```bash
go test ./tests -v
```

O teste irá realizar uma requisição GET para a API, validar o status da resposta, o formato dos dados e se os campos obrigatórios estão presentes e corretos.


## Feito com 💙 por:

- Github [@jussaragranja](https://github.com/jussaragranja)
- Linkedin [@jussaragranja](https://www.linkedin.com/in/jussaragranja/)

