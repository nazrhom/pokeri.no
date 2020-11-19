package services

import "github.com/pokeri.no/api"

type SocketService struct {

}

func InitSocketService(config *api.SocketConfig) SocketService {
	return SocketService{}
}

func (service *SocketService) Write(message string){

}