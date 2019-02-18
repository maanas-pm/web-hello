package main

import (
	"fmt"
	"net/http"
	"github.com/spf13/viper"
)
func init() {
        viper.SetConfigFile(`config.json`)
        err := viper.ReadInConfig()

        if err != nil {
                panic(err)
        }

        if viper.GetBool(`debug`) {
                fmt.Println("Service RUN on DEBUG mode")
        }

}
func hello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
switch r.Method {
    case "GET":
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    case "POST":
	if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8082", nil)
}
