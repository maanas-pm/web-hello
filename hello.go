package main

import (
	"fmt"
	"net/http"
	"github.com/spf13/viper"
        "log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"context"
	"time"
	"go.etcd.io/etcd/client"
	
	_controller_logs "github.com/maanas-pm/web-hello/controller/logs"
)

const (
	LOG_LEVEL = "/sample-app/log_level"
)
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/logs", _controller_logs.Routes())
	})

	return router
}

func init() {
	fmt.Println("inside init")
        viper.SetConfigFile(`config/config.json`)
        err := viper.ReadInConfig()

        if err != nil {
                panic(err)
        }

        if viper.GetBool(`debug`) {
                fmt.Println("Service RUN on DEBUG mode")
        }
	
	var etcd_url = viper.GetString(`etcd.address`) + ":" + viper.GetString(`etcd.port`)
	cfg := client.Config{
		Endpoints:               []string{etcd_url},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	log.Print("Setting '"+ LOG_LEVEL +"' key with '"+ viper.GetString("debug") +"' value")
	resp, err := kapi.Set(context.Background(), LOG_LEVEL, viper.GetString("debug"), nil)
	if err != nil {
		log.Println(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// get "/sample-app/log_level" key's value
	log.Print("Getting '"+ LOG_LEVEL +"' key value")
	resp, err = kapi.Get(context.Background(), LOG_LEVEL, nil)
	if err != nil {
		log.Println(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

func main() {
	//http.HandleFunc("/", hello)
	//http.ListenAndServe(":8082", nil)
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Fatal(http.ListenAndServe(":8082", router)) // Note, the port is usually gotten from the environment.
}
