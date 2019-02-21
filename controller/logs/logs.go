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
	router.Get("/{todoID}", GetALog)
	router.Delete("/{todoID}", DeleteLog)
	router.Post("/", AddLog)
	router.Get("/", GetAllLogs)
	return router
}

func GetALog(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	val, ok := m[todoID]
	if (ok){
		render.JSON(w, r, todos)
	}
	else{
		render.JSON(w, r, {"message": "Requested log not found"})
	}
}

func DeleteLog(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response) // Return some demo response
}

func AddLog(w http.ResponseWriter, r *http.Request) {
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
	val, ok := m[t.Id]
	if(ok){
		t.Time = time.Now()
		t.Request = req
		m[t.Id] = t
		render.JSON(w, r, t)
	}
	else{
		render.JSON(w, r, {"message": "Log Id already exists, cannot override logs"})
	}
}

func GetAllLogs(w http.ResponseWriter, r *http.Request) {
	t := make([]models.Log, len(m))
	for _, v := range m { 
		t = append(t,v)
	}
	render.JSON(w, r, t) // A chi router helper for serializing and returning json
}
