package main

import (
	"fmt"
	"net/http"
	"github.com/spf13/viper"
        "log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	//"context"
	//"time"
	//"go.etcd.io/etcd/client"
	
	_controller_test "github.com/maanas-pm/web-hello/controller/test"
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
		r.Mount("/api/test", _controller_test.Routes())
	})

	return router
}

func init() {
        viper.SetConfigFile(`config/config.json`)
        err := viper.ReadInConfig()

        if err != nil {
                panic(err)
        }

        if viper.GetBool(`debug`) {
                fmt.Println("Service RUN on DEBUG mode")
        }
	
	var etcd_url = viper.GetString(`etcd.address`) + ":" + viper.GetString(`etcd.port`)
	fmt.Println(etcd_url)
	/*cfg := client.Config{
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
	// set "/foo" key with "bar" value
	log.Print("Setting '/foo' key with 'bar' value")
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// get "/foo" key's value
	log.Print("Getting '/foo' key value")
	resp, err = kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
        log.Print("Setting '/foo' key with 'kafkastream_test' value")
        resp, err = kapi.Set(context.Background(), "/foo", "kafkastream_test", nil)
        if err != nil {
                log.Fatal(err)
        } else {
                // print common key info
                log.Printf("Set is done. Metadata is %q\n", resp)
        }
	log.Print("Getting '/foo' key value")
        resp, err = kapi.Get(context.Background(), "/foo", nil)
        if err != nil {
                log.Fatal(err)
        } else {
                // print common key info
                log.Printf("Get is done. Metadata is %q\n", resp)
                // print value
                log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
        }*/
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
