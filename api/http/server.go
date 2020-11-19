package http

import (
	"net/http"
)

func initControllers() []Controller {
	controllers := make([]Controller, 0, 0)
	gameController := new(StartGameController)
	actionController := new(ActionController)
	controllers = append(controllers, gameController)
	controllers = append(controllers, actionController)
	return controllers
}

func initServer() {
	controllers := initControllers()
	for _, c := range controllers {
		http.HandleFunc(c.Route(), c.Handle)
	}
}
func main() {

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
