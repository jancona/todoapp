# todoapp

A sample client server ToDo app based on http://todomvc.com/ and https://www.todobackend.com/

## Running it

You need a [Go 1.14+ installation](https://golang.org/dl/). After cloning this repo, cd into the `server/` directory and type `go run server.go`. The point your browser at:

* `http://localhost:3000/flutter/` for the Flutter UI 
* `http://localhost:3000/vecty/` for the Vecty WASM UI 
* `http://localhost:3000/swagger/` for the Swagger API documentation

## Goal and Approach

The goal of this project is to evaluate and familiarize myself with a variety of tools for building web apps. Some of the things I think I know:

* Want to build backend in Go.
* Would like to do API-first development.
* Would like to use Swagger/OpenAPI for specifying and documenting the API.
* Might want to use Swagger/OpenAPI for automating parts of development process, i.e. code generation.
* [Flutter](https://flutter.dev/) is an interesting tool for building front-ends.

I decided to proceed by building a version of the [ToDoMVC app](http://todomvc.com/) with the tools I wanted to evaluate. Because I want a server backend with an API, I decided to implement the API used by [todobackend.com](https://www.todobackend.com/). They maintain a [test suite](https://www.todobackend.com/specs/index.html) for a standard backend. I also found some Swagger definitions for a ToDo backend, although none of them exactly matched the todobackend.com spec.

I first experimented with a workflow using a Swagger/OpenAPI defintion to generate a Go server framework. All of the tools I tried had weaknesses, either in the Go code they generated or the workflow to use them.

I decided instead to author the Go server by hand, using `net/http` and Gorilla Mux. I used [Swag](https://github.com/swaggo/swag) to generate Swagger and associated documentation pages. It was pretty easy to add annotations in comments and use `go:generate` to run Swag to keep doc in sync with code.

I experimented with two different front ends. I used Flutter to build a (so far) incomplete one. Because I had some experience with [Vecty](https://github.com/gopherjs/vecty) a Go/WASM front end framework, I took their ToDoMVC [example](https://github.com/gopherjs/vecty/tree/master/example/todomvc) and hooked it up to the server. For both front ends, I used the [openapi-generator tool](https://openapi-generator.tech/) to generate client code.

