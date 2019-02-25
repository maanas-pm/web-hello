package controller

import(
	"fmt"
	"errors"
	"bytes"
	"net/http"
	"encoding/json"
    	"net/http/httptest"
	"testing"
	"io/ioutil"
	"github.com/maanas-pm/web-hello/models"
)

const(
	INVALID_JSON = "json: cannot unmarshal string into Go value of type models.Log"
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
        if response.Code != 400{
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
        }
        if obj["message"] != ID_ALREADY_EXISTS{
                t.Error("Output doesn't match, Expected : "+ID_ALREADY_EXISTS +", Got : "+obj["message"])
        }
}

func TestAddLog_failure_invalid_body(t *testing.T){
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
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
        }
	if obj["message"] != INVALID_JSON{
                t.Error("Output doesn't match, Expected : "+ INVALID_JSON +" , Got : "+obj["message"])
        }
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestAddLog_failure_no_body(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        req, err := http.NewRequest("POST", "/",errReader(0))
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
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
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
        if response.Code != 400{
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
        }
        if obj["message"] != ID_NOT_FOUND{
                t.Error("Output doesn't match, Expected : "+ ID_NOT_FOUND + " , Got : "+obj["message"])
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
        if response.Code != 400{
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
        }
        if obj["message"] != ID_NOT_INT_FORMAT{
                t.Error("Output doesn't match, Expected : "+ ID_NOT_INT_FORMAT +" , Got : "+obj["message"])
        }
}

func TestDeleteLog_fail_invalid_id(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        req, err := http.NewRequest("DELETE", "/abc", nil)
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
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
        }
        if obj["message"] != ID_NOT_INT_FORMAT{
                t.Error("Output doesn't match, Expected : "+ ID_NOT_INT_FORMAT +" , Got : "+obj["message"])
        }
}

func TestDeleteLog_fail(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")
        }
        req, err := http.NewRequest("DELETE", "/2", nil)
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
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
        }
        if obj["message"] != ID_NOT_FOUND{
                t.Error("Output doesn't match, Expected : "+ ID_NOT_FOUND +" , Got : "+obj["message"])
        }
}

func TestDeleteLog_success(t *testing.T){
        r:= Routes()
        if r == nil{
                t.Error("initialization failed")   
        }
        req, err := http.NewRequest("DELETE", "/1", nil)
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
                t.Error("Failed: Expected bad request, got "+ string(response.Code))
        }
        if obj["message"] != ID_DELETED{
                t.Error("Output doesn't match, Expected : "+ ID_DELETED +" , Got : "+obj["message"])
        }
}
