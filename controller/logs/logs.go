package controller

import (
	"net/http"
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

var m map[int64]models.Log

func Routes() *chi.Mux {
	if m == nil{
		m = make(map[int64]models.Log)
	}
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
	req := r.Method + " " + r.Host + " "+ r.URL.Path
	body, err := ioutil.ReadAll(r.Body)
    	if err != nil {
        	panic(err)
    	}
	var t models.Log
	err = json.Unmarshal(body, &t)
    	if err != nil {
        	panic(err)
    	}
	t.Time = time.Now()
	t.Request = req
	m[t.Id] = t
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
	render.JSON(w, r, m) // A chi router helper for serializing and returning json
}
