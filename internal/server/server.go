package server

import "github.com/gorilla/mux"

type server struct {
	router *mux.Router
	//logger       *logrus.Logger
	//store        store.Store
	//sessionStore sessions.Store
}

func newServer() *server {
	s := &server{
		mux.NewRouter(),
	}

	return s
}

func (s *server) configureRouter() {

}
