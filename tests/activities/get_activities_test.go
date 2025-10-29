package main

import (
	"encoding/json" // Serializar/Desserializar dados JSON
	"testing"       // Pacote padrão do Go para testes
	"time"          // Manipular e validar datas

	"github.com/go-resty/resty/v2"       // Requisições HTTP
	"github.com/stretchr/testify/assert" // Validações com assert
)

// Criando o contrato da resposta
type Activity struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	DueDate   string `json:"dueDate"`
	Completed bool   `json:"completed"`
}

func TestGetActivities(t *testing.T) {
	client := resty.New() // Criando cliente HTTP com a lib resty
	resp, err := client.R(). // Resultado da requisição salvo em resp, e erros em err
		SetHeader("Accept", "text/plain; v=1.0").
		Get("https://fakerestapi.azurewebsites.net/api/v1/Activities") // Montando e executando a requisição GET

	assert.NoError(t, err)                                                    // Verifica se não houve erro ao fazer a requisição. Se houver, o teste falha aqui.
	assert.Equal(t, 200, resp.StatusCode())                                   // Garante que a resposta HTTP tem status 200 (OK). Se não for, o teste falha.
	assert.Contains(t, resp.Header().Get("Content-Type"), "application/json") // Verifica se o header Content-Type da resposta contém "application/json", ou seja, se a resposta está em formato JSON.

	var activities []Activity                                                       // Criando variavel activities que é array de Activity
	err = json.Unmarshal(resp.Body(), &activities)                                  // Converte corpo da resposta de JSON para array de structs
	assert.NoError(t, err, "Resposta deve ser um array de Activity")                // Se a conversão falhar o teste falha aqui
	assert.Greater(t, len(activities), 0, "Deve retornar pelo menos uma atividade") // Garante que a resposta tem pelo menos um item no array.

	first := activities[0]                           // Validando primeiro item da resposta, primeiro do array
	assert.NotZero(t, first.Id)                      // valida que campo Id não é Zero
	assert.NotEmpty(t, first.Title)                  // valida que campo Title não está vazio
	assert.NotEmpty(t, first.DueDate)                // valida que campo DueDate não está vazio
	_, err = time.Parse(time.RFC3339, first.DueDate) // valida que campo DueDate esta no formato de data RFC3339 (padrão ISO8601)
	assert.NoError(t, err, "DueDate campo data deve estar em formato RFC3339")

	// Validando contrato de todos os itens do array
	// Para cada item retornado pela API:
	for _, act := range activities {
		assert.NotZero(t, act.Id)                       // valida que o Id não é zero
		assert.NotEmpty(t, act.Title)                   // valida que o Title não está vazio
		assert.NotEmpty(t, act.DueDate)                 // valida que o DueDate não está vazio
		_, err := time.Parse(time.RFC3339, act.DueDate) // valida que campo DueDate esta no formato de data RFC3339 (padrão ISO8601)
		assert.NoError(t, err, "DueDate campo data deve estar em formato RFC3339")
	}
}
