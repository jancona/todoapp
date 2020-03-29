package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/jancona/todoapp/server/model"
	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	app := App{
		ToDos: make(map[uuid.UUID]model.ToDo),
	}
	r := NewRouter(app)

	request, _ := http.NewRequest("GET", "/todos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	var empty []model.ToDo
	err := json.NewDecoder(response.Body).Decode(&empty)
	if err != nil {
		t.Errorf("Error parsing JSON: %v", err)
	}
	assert.Equal(t, 0, len(empty), "Expected empty slice, got %#v", empty)
	assert.Equal(t,
		contentType,
		response.Result().Header["Content-Type"][0])
	// Add a ToDo
	tdi := model.ToDoInput{
		Title: "First",
		Order: 1,
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err = enc.Encode(tdi)
	if err != nil {
		t.Errorf("Error encoding ToDoInput: %v", err)
	}
	// missing MIME type
	request, _ = http.NewRequest("POST", "/todos", buf)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnsupportedMediaType, response.Code, "415 response is expected")
	assert.Contains(t, response.Result().Header, "Content-Type", "Should have Content-Type header")
	assert.Equal(t,
		contentType,
		response.Result().Header["Content-Type"][0])

	buf = new(bytes.Buffer)
	enc = json.NewEncoder(buf)
	err = enc.Encode(tdi)
	if err != nil {
		t.Errorf("Error encoding ToDoInput: %v", err)
	}
	// wrong MIME type
	request, _ = http.NewRequest("POST", "/todos", buf)
	request.Header.Add("Content-Type", "application/notjson")
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnsupportedMediaType, response.Code)
	assert.Contains(t, response.Result().Header, "Content-Type", "Should have Content-Type header")
	assert.Equal(t,
		contentType,
		response.Result().Header["Content-Type"][0])

	buf = new(bytes.Buffer)
	enc = json.NewEncoder(buf)
	err = enc.Encode(tdi)
	if err != nil {
		t.Errorf("Error encoding ToDoInput: %v", err)
	}
	// correct MIME type
	request, _ = http.NewRequest("POST", "/todos", buf)
	request.Header.Add("Content-Type", contentType)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, response.Result().Header, "Content-Type", "Should have Content-Type header")
	assert.Equal(t,
		contentType,
		response.Result().Header["Content-Type"][0])
	var created model.ToDo
	err = json.NewDecoder(response.Body).Decode(&created)
	if err != nil {
		t.Errorf("Error parsing JSON: %v", err)
	}
	assert.Equal(t, tdi.Title, created.Title, "Expected Title to match")
	assert.Equal(t, tdi.Order, created.Order, "Expected Order to match")
	assert.Equal(t, tdi.Completed, created.Completed, "Expected Completed to match")
	assert.NotEmpty(t, created.ID)
	t.Logf("app.ToDos: %#v", app.ToDos)
	// GET /todos should now return the created ToDo
	request, _ = http.NewRequest("GET", "/todos", nil)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t,
		contentType,
		response.Result().Header["Content-Type"][0])
	var ret []model.ToDo
	err = json.NewDecoder(response.Body).Decode(&ret)
	if err != nil {
		t.Errorf("Error parsing JSON: %v", err)
	}
	assert.Equal(t, 1, len(ret))
	assert.Equal(t, created, ret[0])
	// GET /todos/{id} should now return the created ToDo
	request, _ = http.NewRequest("GET", "/todos/"+created.ID.String(), nil)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Contains(t, response.Result().Header, "Content-Type", "Should have Content-Type header")
	assert.Equal(t,
		contentType,
		response.Result().Header["Content-Type"][0])
	var ret2 model.ToDo
	err = json.NewDecoder(response.Body).Decode(&ret2)
	if err != nil {
		t.Errorf("Error parsing JSON: %v", err)
	}
	assert.Equal(t, created, ret2)
	assert.Contains(t, response.Result().Header, "Content-Type", "Should have Content-Type header")
	assert.Equal(t,
		contentType,
		response.Result().Header["Content-Type"][0])
}

// func TestGetTodo(t *testing.T) {
// }
// func TestPostTodo(t *testing.T) {
// }
// func TestPutTodo(t *testing.T) {
// }
// func TestDeleteTodo(t *testing.T) {
// }
