package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Server : server structure
type Server struct {
	router   *mux.Router
	Exit     chan int
	port     string
	shutdown string
}

//NewServer : method to create new server object
func NewServer(port string, exit chan int) (Server, error) {
	out := Server{router: mux.NewRouter(), Exit: exit, port: port}
	if exit == nil || len(port) < 2 {
		return out, errors.New("Internal exception")
	}
	return out, nil
}

func (s *Server) routes() {
	http.Handle("/", s.router)
	s.router.HandleFunc("/nsd", calculateNSD).Methods("POST")

}

type shutdownAuth struct {
	Token string `json:"token,omitempty"`
}

//Start : method to execute server
func (s *Server) Start(key string) {
	//s.routes()
	fmt.Println("Server Started on port " + s.port)
	s.shutdown = key
	http.Handle("/", s.router)
	s.router.HandleFunc("/gcd", calculateNSD).Methods("POST")
	s.router.HandleFunc("/abcd", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var resp source
		resp.Things = primes
		out, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("shit shit shit")
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Fprintf(w, (string) out)

	})
	s.router.HandleFunc("/hello", sayHello)
	s.router.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		var auth shutdownAuth
		_ = json.NewDecoder(r.Body).Decode(&auth)
		if auth.Token == key {
			s.Exit <- 0
		}
	}).Methods("POST")
	http.ListenAndServe(s.port, s.router)
}

func notImplemented() {
	fmt.Printf("Not implemented")
}
