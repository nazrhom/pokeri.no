package http

import (
	log "github.com/sirupsen/logrus"
	netHttp "net/http"
)

type Controller interface {
	Route() string
	Handle(w netHttp.ResponseWriter, req *netHttp.Request)
}
type StartGameController struct {
}

func (controller *StartGameController) Handle(w netHttp.ResponseWriter, req *netHttp.Request) {
	log.Info("Handling Start Game Request")
}

func (controller *StartGameController) Route() string {
	return "/game"
}

type ActionController struct {
}

func (controller *ActionController) Handle(w netHttp.ResponseWriter, req *netHttp.Request) {
	log.Info("Handling Action")
}

func (controller *ActionController) Route() string {
	return "/action"
}
