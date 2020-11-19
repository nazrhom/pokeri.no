package http

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)


func initServer() {
	gameController := new(GameController)
	http.HandleFunc("/game", gameController.StartGame)
	http.HandleFunc("/action", gameController.Action)

}
func Start() {
	initServer()
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Error(fmt.Sprintf("Error %v caught while starting server, exiting...", err))
		panic(err)
	}
}
