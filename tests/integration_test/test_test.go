package integrationtest

import (
	"answerService/internals/dto"
	"answerService/internals/models"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAnswer_Integration(t *testing.T) {
	app := setupTestApp(t)

	question := models.Question{Text: "What is Go?"}
	_, err := app.Service.SrvQuestion.CreateQuestion(context.Background(), &question)
	assert.NoError(t, err)
	assert.Equal(t, 2, question.ID)

	body := dto.AnswerCreateRequest{
		Text:   "It is a language",
		UserID: "user-123",
	}
	b, err := json.Marshal(body)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/questions/"+strconv.Itoa(int(question.ID))+"/answers", bytes.NewBuffer(b))

	w := httptest.NewRecorder()
	app.Server.CreateAnswer(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var created models.Answer
	err = json.Unmarshal(w.Body.Bytes(), &created)
	assert.NoError(t, err)
	assert.Equal(t, int(question.ID), created.QuestionID)
	assert.Equal(t, body.Text, created.Text)
	assert.Equal(t, body.UserID, created.UserID)

	var dbAnswer *models.Answer
	dbAnswer, err = app.Storage.Answer(context.Background(), 2)
	assert.NoError(t, err)
	assert.Equal(t, body.UserID, dbAnswer.UserID)
}
