package todo

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// List - GET /api/todo - get list of todos
func List() APIResponse {
	log.Trace("API List")
	url := getEndpoint(fmt.Sprintf("/api/todo"))
	resp := Request("GET", url, []byte{})
	return resp
}

// Create - POST /api/todo - get list of todos
func Create(todo Todo) APIResponse {
	log.Trace("API Create")
	marshalled, _ := json.Marshal(todo)
	url := getEndpoint(fmt.Sprintf("/api/todo/"))
	resp := Request("POST", url, []byte(marshalled))
	return resp
}

// Get - GET /api/todo/{token} - get a todo
func Get(token string) APIResponse {
	log.Trace("API Get")
	url := getEndpoint(fmt.Sprintf("/api/todo/%v", token))
	resp := Request("GET", url, []byte{})
	return resp
}

// Put - PUT /api/todo - create a new todo
func Put(token string, todo Todo) APIResponse {
	log.Trace("API Update")
	marshalled, _ := json.Marshal(todo)
	url := getEndpoint(fmt.Sprintf("/api/todo/%v", token))
	resp := Request("POST", url, []byte(marshalled))
	return resp
}

// Delete - DELETE /api/todo/{token} - delete a todo
func Delete(token string) APIResponse {
	log.Trace("API Delete")
	url := getEndpoint(fmt.Sprintf("/api/todo/%v", token))
	resp := Request("DELETE", url, []byte{})

	if resp.StatusCode != 200 {
		showFullResponseBody(resp)
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}
