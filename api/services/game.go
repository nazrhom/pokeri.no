package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pokeri.no/api"
	"github.com/pokeri.no/api/persistence"
)

type GameService struct {
	Database *persistence.Database
}

func InitService() GameService {
	db := persistence.InitDatabase(api.DbConfig{
		Hostname: "",
		Port:     "",
		Password: "",
		User:     "",
	})
	service := GameService{
		Database: db,
	}
	return service
}
func (service *GameService) StartGame(players []string, smallBlind int, startingChips int, timer int, buyIn int) (string, error) {
	gameId := uuid.New().String()
	model := api.GameModel{
		Players:       players,
		SB:            smallBlind,
		StartingChips: startingChips,
		BlindsTimer:   timer,
		BuyIn:         buyIn,
		Id:            gameId,
	}
	err := service.Database.SaveGame(model)
	if err != nil {
		return "", err
	}
	return gameId, nil
}

func (service *GameService) LoadGame(gameId string) (api.GameModel, error) {
	game, err := service.Database.LoadGame(gameId)
	if err != nil {
		return api.GameModel{}, err
	}
	return game, nil
}

func (service *GameService) Action(model api.ActionModel) error {
	switch action := model.Name; action {
	case "BET":
		return service.bet(model.Amount)
	case "FOLD":
		return service.fold()
	case "RAISE":
		return service.bet(model.Amount)
	case "CHECK":
		return service.check()
	}
	return errors.New(fmt.Sprintf("Unknown Action %s", model.Name))
}
func (service *GameService) bet(amount float64) error {
	return nil
}

func (service *GameService) check() error {
	return nil
}

func (service *GameService) fold() error {
	return nil
}
