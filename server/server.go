//go:generate swag init -g $GOFILE
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jancona/todoapp/server/docs"
	"github.com/jancona/todoapp/server/model"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	contentType = "application/json"
	defaultURL  = "http://localhost:3000"
	indexHTML   = `<html>
	<body>
		<a href='/flutter/'>Flutter UI</a><br/>
		<a href='/wasm/'>Vecty WASM UI</a><br/>
		<a href='/swagger/'>Swagger API Documentation</a><br/>
		<a href='/todos'>API</a><br/>
	</body>
</html>
`
)

// @title To Do List API
// @version 0.1.0
// @description This is a simple todo list API.

// @contact.name Jim Ancona
// @contact.url https://github.com/jancona
// @contact.email jim@anconafamily.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.ourroots.org
// @BasePath /
// @accept application/json
// @produce application/json
// @schemes http https
func main() {
	// Find out if we're running in a Lambda function
	isLambda := true
	if os.Getenv("LAMBDA_TASK_ROOT") == "" {
		isLambda = false
	}
	// baseURL is used to build proper absolute URLs in a couple of places
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = defaultURL
	}
	log.Printf("BaseURL: %s", baseURL)
	url, err := url.ParseRequestURI(baseURL)
	if err != nil {
		log.Fatalf("Error parsing base URL '%s': %v", baseURL, err)
	}
	log.Printf("BaseURL: %s, url: %s", baseURL, url)
	r := NewRouter(App{
		ToDos:   make(map[uuid.UUID]model.ToDo),
		BaseURL: baseURL,
	})
	docs.SwaggerInfo.Host = url.Hostname()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(baseURL+"/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	r.NotFoundHandler = http.HandlerFunc(notFound)

	if isLambda {
		// Lambda-specific setup
		// Note that the Lamda doesn't serve static content, only the API
		// API Gateway proxies static content requests directly to an S3 bucket
		// API Gateway + Lambda is https-only
		docs.SwaggerInfo.Schemes = []string{"https"}
		adapter := gorillamux.New(r)
		lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			log.Printf("Lambda request %#v", req)
			// If no name is provided in the HTTP request body, throw an error
			return adapter.ProxyWithContext(ctx, req)
		})
		log.Fatal("Lambda exiting...")
	} else {
		// If we're not running in Lambda we also serve the static content.
		// This is useful in development. It might also be in a traditional server deploy, but requirements
		// for all of this are TBD.
		flutterDir := "../flutterui/build/web"
		vectyDir := "../vectyui/build/web"
		log.Printf("flutterDir: %s, vectyDir: %s", flutterDir, vectyDir)
		r.PathPrefix("/flutter/").
			Handler(http.StripPrefix("/flutter", http.FileServer(http.Dir(flutterDir))))
		r.PathPrefix("/wasm/").
			Handler(http.StripPrefix("/wasm", http.FileServer(http.Dir(vectyDir))))
		log.Fatal(http.ListenAndServe(url.Host,
			handlers.LoggingHandler(os.Stdout,
				handlers.CORS(
					handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
					handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
					handlers.AllowedOrigins([]string{"*"}))(r))))
	}
}

// NewRouter builds a router for handling requests
func NewRouter(app App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", app.GetIndex).Methods("GET")
	r.HandleFunc("/index.html", app.GetIndex).Methods("GET")
	r.HandleFunc("/todos", app.GetAllToDos).Methods("GET")
	r.HandleFunc("/todos", app.PostToDo).Methods("POST")
	r.HandleFunc("/todos", app.DeleteAllToDos).Methods("DELETE")
	r.HandleFunc("/todos/{id}", app.GetToDo).Methods("GET")
	r.HandleFunc("/todos/{id}", app.PatchToDo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", app.DeleteToDo).Methods("DELETE")
	return r
}

// App is the container for the application
type App struct {
	BaseURL string
	// Dummy "database".
	// (Note that this behaves really strangely when multiple Lambda are running
	// because which database you see depends on which Lamda instance you're routed to.)
	ToDos map[uuid.UUID]model.ToDo
}

// GetIndex returns an HTML index page
func (app App) GetIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(indexHTML))
}

// GetAllToDos returns all todos in the database
// @summary returns all todos
// @router /todos [get]
// @tags todos
// @id find
// @produce application/json
// @success 200 {array} model.ToDo "OK"
// @failure 500 {object} model.Error "Server error"
func (app App) GetAllToDos(w http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", contentType)
	v := make([]model.ToDo, 0, len(app.ToDos))

	for _, value := range app.ToDos {
		v = append(v, value)
	}
	err := enc.Encode(v)
	if err != nil {
		serverError(w, err)
		return
	}
}

// DeleteAllToDos deletes all ToDos from the database
// @summary deletes all ToDos
// @router /todos [delete]
// @tags todos
// @id deleteAll
// @success 200 {array} model.ToDo "OK"
// @failure 500 {object} model.Error "Server error"
func (app App) DeleteAllToDos(w http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", contentType)

	for id := range app.ToDos {
		delete(app.ToDos, id)
	}
	v := make([]model.ToDo, 0)
	err := enc.Encode(v)
	if err != nil {
		serverError(w, err)
		return
	}
}

// PostToDo adds a new ToDo to the database
// @summary adds a new ToDo
// @router /todos [post]
// @tags todos
// @id addOne
// @Param todo body model.ToDoInput true "Add ToDo"
// @accept application/json
// @produce application/json
// @success 201 {object} model.ToDo "OK"
// @failure 415 {object} model.Error "Bad Content-Type"
// @failure 500 {object} model.Error "Server error"
func (app App) PostToDo(w http.ResponseWriter, req *http.Request) {
	mt, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		msg := fmt.Sprintf("Bad Content-Type '%s'", mt)
		log.Print(msg)
		errorResponse(w, http.StatusUnsupportedMediaType, fmt.Sprintf("Bad MIME type '%s'", mt))
		return
	}
	if mt != contentType {
		msg := fmt.Sprintf("Bad Content-Type '%s'", mt)
		log.Print(msg)
		errorResponse(w, http.StatusUnsupportedMediaType, fmt.Sprintf("Bad MIME type '%s'", mt))
		return
	}
	var todo model.ToDo
	err = json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		msg := fmt.Sprintf("Bad request: %v", err.Error())
		log.Print(msg)
		errorResponse(w, http.StatusBadRequest, msg)
		return
	}
	id := uuid.New()
	todo.ID = &id
	todo.URL = app.BaseURL + "/todos/" + todo.ID.String()
	// Add to "database"
	app.ToDos[*todo.ID] = todo
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	err = enc.Encode(todo)
	if err != nil {
		serverError(w, err)
		return
	}
}

// GetToDo gets a ToDo from the database
// @summary gets a ToDo
// @router /todos/{id} [get]
// @tags todos
// @id getOne
// @Param id path string true "ToDo ID" format(uuid)
// @produce application/json
// @success 200 {object} model.ToDo "OK"
// @failure 404 {object} model.Error "Not found"
// @failure 500 {object} model.Error "Server error"
func (app App) GetToDo(w http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", contentType)
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		notFound(w, req)
		return
	}
	todo, found := app.ToDos[id]
	if !found {
		notFound(w, req)
		return
	}
	err = enc.Encode(todo)
	if err != nil {
		serverError(w, err)
		return
	}
}

// PatchToDo updates a ToDo in the database
// @summary updates a ToDo
// @router /todos/{id} [patch]
// @tags todos
// @id update
// @Param id path string true "ToDo ID" format(uuid)
// @Param todo body model.ToDoInput true "Update ToDo"
// @accept application/json
// @produce application/json
// @success 200 {object} model.ToDo "OK"
// @failure 415 {object} model.Error "Bad Content-Type"
// @failure 500 {object} model.Error "Server error"
func (app App) PatchToDo(w http.ResponseWriter, req *http.Request) {
	mt, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		serverError(w, err)
		return
	}
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		notFound(w, req)
		return
	}
	_, found := app.ToDos[id]
	if !found {
		// Not allowed to add a ToDo with PATCH
		notFound(w, req)
		return
	}
	if mt != contentType {
		msg := fmt.Sprintf("Bad Content-Type '%s'", mt)
		log.Print(msg)
		errorResponse(w, http.StatusUnsupportedMediaType, fmt.Sprintf("Bad MIME type '%s'", mt))
		return
	}
	var tdi model.ToDoInput
	err = json.NewDecoder(req.Body).Decode(&tdi)
	if err != nil {
		msg := fmt.Sprintf("Bad request: %v", err.Error())
		log.Print(msg)
		errorResponse(w, http.StatusBadRequest, msg)
		return
	}
	todo := app.ToDos[id]
	todo.Title = tdi.Title
	todo.Order = tdi.Order
	todo.Completed = tdi.Completed
	// Add to "database"
	app.ToDos[id] = todo
	w.Header().Set("Content-Type", contentType)
	enc := json.NewEncoder(w)
	err = enc.Encode(todo)
	if err != nil {
		serverError(w, err)
		return
	}
}

// DeleteToDo deletes a ToDo from the database
// @summary deletes a ToDo
// @router /todos/{id} [delete]
// @tags todos
// @id delete
// @Param id path string true "ToDo ID" format(uuid)
// @success 204 {object} model.ToDo "OK"
// @failure 500 {object} model.Error "Server error"
func (app App) DeleteToDo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		notFound(w, req)
		return
	}
	delete(app.ToDos, id)
	w.WriteHeader(http.StatusNoContent)
}

func serverError(w http.ResponseWriter, err error) {
	log.Print("Server error: " + err.Error())
	// debug.PrintStack()
	errorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", err.Error()))
}

func notFound(w http.ResponseWriter, req *http.Request) {
	m := fmt.Sprintf("Path '%s' not found", req.URL.RequestURI())
	log.Print(m)
	errorResponse(w, http.StatusNotFound, m)
}

func errorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	e := model.Error{Code: code, Message: message}
	err := enc.Encode(e)
	if err != nil {
		log.Printf("Failure encoding error response: '%v'", err)
	}
}
