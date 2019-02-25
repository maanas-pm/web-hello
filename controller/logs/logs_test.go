package controller

import(
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
    	"net/http/httptest"
	"testing"
	"io/ioutil"
	"github.com/maanas-pm/web-hello/models"
)

func TestRoutes(t *testing.T){
	r:= Routes()
	if r == nil{
		t.Error("initialization failed")
	}
}

func TestAddLog_success(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
	var log_obj models.Log
        log_obj.Id = 1
	log_obj_json, _ := json.Marshal(log_obj)
        req, err := http.NewRequest("POST", "/", bytes.NewBuffer(log_obj_json))
        if err != nil {
                t.Fatal(err)
        }
        response := httptest.NewRecorder()
        r.ServeHTTP(response, req)

        if err != nil {
                panic(err.Error())
        }
        body, e := ioutil.ReadAll(response.Body)
        if e != nil {
                panic(e)
        }
	var obj models.Log
        json.Unmarshal(body, &obj)
        fmt.Println(obj)
        if response.Code != 200{
                t.Error("Failed to get logs")
        }
        if obj.Id != log_obj.Id{
                t.Error("Output doesn't match, Expected : "+ string(log_obj.Id)+", Got : "+string(obj.Id))
        }
}

func TestAddLog_re_add(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        var log_obj models.Log
        log_obj.Id = 1
        log_obj_json, _ := json.Marshal(log_obj)
        req, err := http.NewRequest("POST", "/", bytes.NewBuffer(log_obj_json))
        if err != nil {
                t.Fatal(err)
        }
        response := httptest.NewRecorder()
        r.ServeHTTP(response, req)

        if err != nil {
                panic(err.Error())
        }
        body, e := ioutil.ReadAll(response.Body)
        if e != nil {
                panic(e)
        }
        obj := make(map[string]string)
        json.Unmarshal(body, &obj)
        fmt.Println(obj)
        if response.Code != 200{
                t.Error("Failed to get logs")
        }
        if obj["message"] != "Log Id already exists, cannot override logs"{
                t.Error("Output doesn't match, Expected : Log Id already exists, cannot override logs, Got : "+obj["message"])
        }
}

func TestAddLog_failure(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        log_obj_json, _ := json.Marshal("{}")
        req, err := http.NewRequest("POST", "/", bytes.NewBuffer(log_obj_json))
        if err != nil {
                t.Fatal(err)
        }
        response := httptest.NewRecorder()
        r.ServeHTTP(response, req)

        if err != nil {
                panic(err.Error())
        }
        body, e := ioutil.ReadAll(response.Body)
        if e != nil {
                panic(e)
        }
        obj := make(map[string]string)
        json.Unmarshal(body, &obj)
        fmt.Println(obj)
        if response.Code != 400{
                t.Error("Failed: request processed")
        }
}

func TestGetAllLog(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        req, err := http.NewRequest("GET", "/", nil)
        if err != nil {
                t.Fatal(err)
        }
        response := httptest.NewRecorder()
        r.ServeHTTP(response, req)
        if response.Code != 200{
                t.Error("Failed to get logs")
        }
}

func TestGetALog(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        req, err := http.NewRequest("GET", "/1", nil)
        if err != nil {
                t.Fatal(err)
        }
        response := httptest.NewRecorder()
        r.ServeHTTP(response, req)

        if err != nil {
                panic(err.Error())
        }
        body, e := ioutil.ReadAll(response.Body)
        if e != nil {
                panic(e)
        }
        var obj models.Log
        json.Unmarshal(body, &obj)
        fmt.Println(obj)
        if response.Code != 200{
                t.Error("Failed to get logs")
        }
        if obj.Id != 1{
                t.Error("Output doesn't match, Expected : Requested Log with Id 1 , Got : "+ string(obj.Id))
        }
}

func TestGetALog_Not_Found(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        req, err := http.NewRequest("GET", "/2", nil)
        if err != nil {
                t.Fatal(err)
        }
        response := httptest.NewRecorder()
        r.ServeHTTP(response, req)

        if err != nil {
                panic(err.Error())
        }
        body, e := ioutil.ReadAll(response.Body)
        if e != nil {
                panic(e)
        }
        obj := make(map[string]string)
        json.Unmarshal(body, &obj)
        fmt.Println(obj)
        if response.Code != 200{
                t.Error("Failed to get logs")
        }
        if obj["message"] != "Requested log not found"{
                t.Error("Output doesn't match, Expected : Requested log not found, Got : "+obj["message"])
        }
}

func TestGetALog_fail(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        req, err := http.NewRequest("GET", "/abc", nil)
        if err != nil {
                t.Fatal(err)
        }
        response := httptest.NewRecorder()
        r.ServeHTTP(response, req)

        if err != nil {
                panic(err.Error())
        }
        body, e := ioutil.ReadAll(response.Body)
        if e != nil {
                panic(e)
        }
        obj := make(map[string]string)
        json.Unmarshal(body, &obj)
        fmt.Println(obj)
        if response.Code != 200{
                t.Error("Failed to get logs")
        }
        if obj["message"] != "Requested log id is not in int format"{
                t.Error("Output doesn't match, Expected : Requested log id is not in int format, Got : "+obj["message"])
        }
}
