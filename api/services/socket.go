package services

import "github.com/pokeri.no/api"

type GameStartedMessage struct {
	GameStatus api.GameModel `json:"game_status"`
}

type HandStartedMessage struct {
	GameStatus api.GameModel `json:"game_status"`
}

type HandEndedMessage struct {
	GameStatus api.GameModel `json:"game_status"`
}

type FlopServedMessage struct {
	Cards []api.CardModel `json:"cards"`
}
type TurnServedMessage struct {
	Card api.CardModel `json:"card"`
}

type RiverServedMessage struct {
	Card api.CardModel `json:"card"`
}

type SocketService struct {
}

func InitSocketService(config *api.SocketConfig) SocketService {
	return SocketService{}
}

func (service *SocketService) Write(message string) {

}
