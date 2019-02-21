package controller

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"log"
	"time"
 	"io/ioutil"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"encoding/json"
	"github.com/maanas-pm/web-hello/models"
)

type Todo struct {
	Id  string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo)
	router.Delete("/{todoID}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Id:  todoID,
		Title: "Hello world",
		Body:  "Heloo world from planet earth",
	}
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response) // Return some demo response
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, string(requestDump))
	}
	log.Println(requestDump)
	req := r.Method + " " + r.Host + " "+ string(r.URL)
	body, err := ioutil.ReadAll(r.Body)
    	if err != nil {
        	panic(err)
    	}
    	log.Println(string(body))
	var t models.Log
	err = json.Unmarshal(body, &t)
    	if err != nil {
        	panic(err)
    	}
	t.Time = time.Now()
	t.Request = string(req)
	render.JSON(w, r, t) // Return some demo response
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Id:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}