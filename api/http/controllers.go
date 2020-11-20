package http

import (
	"encoding/json"
	"github.com/pokeri.no/api"
	"github.com/pokeri.no/api/services"
	log "github.com/sirupsen/logrus"
	"io"
	netHttp "net/http"
)

type Controller interface {
	Route() string
	Handle(w netHttp.ResponseWriter, req *netHttp.Request)
}
type GameController struct {
	GameService   *services.GameService
	SocketService *services.SocketService
}

func (controller *GameController) StartGame(w netHttp.ResponseWriter, req *netHttp.Request) {
	log.Info("Handling Start Game Request")
	game, err := DeserializeGameRequest(req.Body)
	if err != nil {
		netHttp.Error(w, err.Error(), netHttp.StatusBadRequest)
		return
	}
	gId, gErr := controller.GameService.StartGame(game.Players, game.SB, game.StartingChips, game.BlindsTimer, game.BuyIn)
	if gErr != nil {
		netHttp.Error(w, gErr.Error(), netHttp.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

func (controller *GameController) LoadGame(w netHttp.ResponseWriter, req *netHttp.Request) {
	log.Info("Handling Load Game Request")
}

func (controller *GameController) StartHand(w netHttp.ResponseWriter, req *netHttp.Request) {
	log.Info("Handling Start Hand")
}
func (controller *GameController) Action(w netHttp.ResponseWriter, req *netHttp.Request) {
	log.Info("Handling Action")
	action, err := DeserializeActionRequest(req.Body)
	if err != nil {
		netHttp.Error(w, err.Error(), netHttp.StatusBadRequest)
		return
	}
	aErr := controller.GameService.Action(action)
	if aErr != nil {
		netHttp.Error(w, aErr.Error(), netHttp.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)

}

func DeserializeActionRequest(body io.ReadCloser) (api.ActionModel, error) {
	var action api.ActionModel
	err := json.NewDecoder(body).Decode(&action)
	if err != nil {
		return api.ActionModel{}, err
	}
	return action, nil
}
func DeserializeGameRequest(body io.ReadCloser) (api.GameModel, error) {
	var game api.GameModel
	err := json.NewDecoder(body).Decode(&game)
	if err != nil {
		return api.GameModel{}, err
	}
	return game, nil
}
