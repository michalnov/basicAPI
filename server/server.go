package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	Exit     chan int
	port     string
	shutdown string
}

func NewServer(port string, exit chan int) (error, Server) {
	out := Server{router: mux.NewRouter(), Exit: exit, port: port}
	if exit == nil || len(port) < 2 {
		return errors.New("Internal exception"), out
	}
	return nil, out
}

func (s *Server) routes() {
	http.Handle("/", s.router)
	s.router.HandleFunc("/nsd", calculateNSD).Methods("GET")

}

type shutdownAuth struct {
	token string `json:"token"`
}

func (s *Server) Start() {
	//s.routes()
	fmt.Println("Server Started on port " + s.port)
	s.shutdown = ""
	http.Handle("/", s.router)
	s.router.HandleFunc("/gcd", calculateNSD).Methods("POST")
	s.router.HandleFunc("/hello", sayHello)
	s.router.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		var auth shutdownAuth
		_ = json.NewDecoder(r.Body).Decode(&auth)
		if auth.token == s.shutdown {
			s.Exit <- 0
		}

	}).Methods("POST")
	http.ListenAndServe(s.port, s.router)
}

func notImplemented() {
	fmt.Printf("Not implemented")
}

/*fmt.Println("Run server on port: " + s.conf.port)
http.Handle("/", s.router)
s.router.HandleFunc("/hello", handler.Hello_Handler).Methods("GET")
s.router.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
	s.Terminate <- 0
})
http.ListenAndServe(s.conf.port, s.router)
*/
