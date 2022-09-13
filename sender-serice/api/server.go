package server

import (
	"log"
	"net/http"

	"github.com/DarkJediDJ/support-chat/sender-serice/api/sender"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

type App struct {
	Router *mux.Router
}

func NewApp() *App {
	return &App{}
}

func (a *App) InitRouter(conn *kafka.Conn) {
	myRouter := mux.NewRouter().StrictSlash(false)
	myRouter.HandleFunc("/", sender.Init(conn).Send).Methods("POST")
	a.Router = myRouter
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
