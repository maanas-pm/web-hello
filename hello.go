package main

import (
	"fmt"
	"net/http"
	"github.com/spf13/viper"
        "log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	
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
	fmt.PrintLn(viper.GetString("etcd.address")
        fmt.PrintLn(viper.GetString("etcd.port")
	if (viper.GetBool(`etcd`) && viper.GetBool(`etcd.address`) && viper.GetBool(`etcd.port`)){
		var etcd_url = viper.GetString(`etcd.address`) + viper.GetString(`etcd.port`)
		fmt.Println(etcd_url)
	}

}
/*func hello(w http.ResponseWriter, r *http.Request){
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
*/
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
