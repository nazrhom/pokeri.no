package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pokeri.no/api"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	client *redis.Client
	ctx    context.Context
}

func InitDatabase(config api.DbConfig) *Database {
	dbAddr := config.Hostname + ":" + config.Port
	client := redis.NewClient(&redis.Options{Addr: dbAddr,
		Password: config.Password,
		DB:       0})
	return &Database{client: client}
}

func (database *Database) SaveGame(model api.GameModel) error {
	serialized, se := Serialize(model)
	if se != nil {
		log.Error(fmt.Sprintf("Error %v caught in serialization of model %v: ", se, model))
	}
	err := database.client.Set(database.ctx, model.Id, serialized, 0).Err()
	if err != nil {
		log.Error(fmt.Sprintf("Error caught while saving game: %v", err))
		return err
	}
	return nil

}
func (database *Database) LoadGame(gameId string) (api.GameModel, error) {
	serialized, err := database.client.Get(database.ctx, gameId).Result()
	if err != nil {
		log.Error(fmt.Sprintf("Error while loading game %s", gameId))
		return api.GameModel{}, err
	}
	model, err := Deserialize(serialized)
	if err != nil {
		log.Error(fmt.Sprintf("Error while deserializing game %s", gameId))
		return api.GameModel{}, err
	}
	return model.(api.GameModel), nil

}
func Serialize(model api.Model) (string, error) {
	bytes, err := json.Marshal(model)
	if err != nil {
		log.Error(fmt.Sprintf("Error %v caught while in json marshaling", err))
		return "", err
	}
	return string(bytes), nil
}

func Deserialize(str string) (api.Model, error) {
	var game api.GameModel
	err := json.Unmarshal([]byte(str), &game)
	if err != nil {
		log.Error("Error while deserializing game")
		return game, err
	}
	return game, nil
}
