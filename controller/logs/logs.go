package controller

import (
	"net/http"
	"time"
 	"io/ioutil"
	"strconv"
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
	router.Get("/{logId}", GetALog)
	router.Delete("/{logId}", DeleteLog)
	router.Post("/", AddLog)
	router.Get("/", GetAllLogs)
	return router
}

func GetALog(w http.ResponseWriter, r *http.Request) {
	logId := chi.URLParam(r, "logId")
	i, err := strconv.ParseInt(logId, 10, 64)
	if err != nil {
    		response := make(map[string]string)
                response["message"] = "Requested log id is not in int format"
                render.JSON(w, r, response)
	}
	val, ok := m[i]
	if ok {
		render.JSON(w, r, val)
	} else {
		response := make(map[string]string)
		response["message"] = "Requested log not found"
		render.JSON(w, r, response)
	}
}

func DeleteLog(w http.ResponseWriter, r *http.Request) {
	logId := chi.URLParam(r, "logId")
        i, err := strconv.ParseInt(logId, 10, 64)
        if err != nil {
                response := make(map[string]string)
                response["message"] = "Requested log id is not in int format"
                render.JSON(w, r, response)
        }
        val, ok := m[i]
        if ok {
                delete(m, logId)
        } else {
                response := make(map[string]string)
                response["message"] = "Requested log not found"
                render.JSON(w, r, response)
        }

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
	_, ok := m[t.Id]
	if ok {
		response := make(map[string]string)
                response["message"] = "Log Id already exists, cannot override logs"
                render.JSON(w, r, response)
	} else {
		t.Time = time.Now()
                t.Request = req
                m[t.Id] = t
                render.JSON(w, r, t)
	}
}

func GetAllLogs(w http.ResponseWriter, r *http.Request) {
	t := make([]models.Log, len(m))
	for k, v := range m { 
		if k > 0 {
			t = append(t,v)
		}
	}
	render.JSON(w, r, t) // A chi router helper for serializing and returning json
}
