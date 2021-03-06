package controller

import (
	"log"
	"net/http"
	"time"
 	"io/ioutil"
	"strconv"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"encoding/json"
	"github.com/maanas-pm/web-hello/models"
)

const (
	ID_NOT_INT_FORMAT = "Requested log id is not in int format"
	ID_NOT_FOUND = "Requested log not found"
	ID_ALREADY_EXISTS = "Log Id already exists, cannot override logs"
	ID_DELETED = "Requested log deleted"
)
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
	response := make(map[string]string)
	logId := chi.URLParam(r, "logId")
	i, err := strconv.ParseInt(logId, 10, 64)
	if err != nil {
                response["message"] = ID_NOT_INT_FORMAT
		render.Status(r, http.StatusBadRequest)
                render.JSON(w, r, response)
	} else {
	val, ok := m[i]
		if ok {
			render.JSON(w, r, val)
		} else {
			response["message"] = ID_NOT_FOUND
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response)
		}
	}
}

func DeleteLog(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	logId := chi.URLParam(r, "logId")
        i, err := strconv.ParseInt(logId, 10, 64)
        if err != nil {
                response := make(map[string]string)
                response["message"] = ID_NOT_INT_FORMAT
		render.Status(r, http.StatusBadRequest)
                render.JSON(w, r, response)
        } else {
        	_, ok := m[i]
        	if ok {
                	delete(m, i)
                	response["message"] = ID_DELETED
                	render.JSON(w, r, response)
        	} else {
                	response["message"] = ID_NOT_FOUND
			render.Status(r, http.StatusBadRequest)
                	render.JSON(w, r, response)
        	}
	}
}

func AddLog(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	req := r.Method + " " + r.Host + " "+ r.URL.Path
	body, err := ioutil.ReadAll(r.Body)
    	if err != nil {
                response["message"] = err.Error()
                render.Status(r, http.StatusBadRequest)
                render.JSON(w, r, response)
	}
	var t models.Log
	err = json.Unmarshal(body, &t)
    	if err != nil {
                response["message"] = err.Error()
		render.Status(r, http.StatusBadRequest) 
                render.JSON(w, r, response)
    	} else {
		_, ok := m[t.Id]
		if ok {
	                response["message"] = ID_ALREADY_EXISTS
        	        render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response)
		} else {
			t.Time = time.Now()
	                t.Request = req
	                m[t.Id] = t
	                render.JSON(w, r, t)
		}
	}
}

func GetAllLogs(w http.ResponseWriter, r *http.Request) {
	t := []models.Log{}
	log.Println(t)
	for k, v := range m {
		log.Println("key : "+string(k)) 
		if k > 0 {
			log.Println(v)
			t = append(t,v)
		}
	}
	render.JSON(w, r, t) // A chi router helper for serializing and returning json
}
