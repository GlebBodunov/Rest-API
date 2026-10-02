package main

import (
	"fmt"
	"restapi/http"
	"restapi/todo"
)

func main() {

	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("fail to start HTTP Server:", err)
	}

}
