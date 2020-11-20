package api

type PhaseName string
type ActionName string

type Model interface{}

type ActionModel struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	GameId string  `json:"game_id"`
}

type CardModel struct {
	Suit   string `json:"suit"`
	Number string `json:"number"`
}

type PlayerStatusModel struct {
	Cards []*CardModel `json:"cards"`
	Stack float64     `json:"stack"`
	InPot float64     `json:"in_pot"`
}

type GameStatusModel struct {
	ButtonPlayer int                 `json:"button_player"`
	CurrPlayer   int                 `json:"curr_player"`
	Players      []*PlayerStatusModel `json:"players"`
	CurrPhase    string              `json:"curr_phase"`
}

type GameModel struct {
	Players       []string  `json:"players"`
	SB            int       `json:"small_blind"`
	StartingChips int       `json:"starting_chips"`
	BlindsTimer   int       `json:"blinds_timer"`
	BuyIn         int       `json:"buy_in"`

	Status *GameStatusModel `json:"status"`
	Id    string      `json:"game_id"`
}
