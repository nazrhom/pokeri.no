package api

type PhaseName string
type ActionName string

const (
	PRE_FLOP PhaseName = "PreFlop"
	FLOP     PhaseName = "Flop"
	TURN     PhaseName = "Turn"
	RIVER    PhaseName = "River"
)

const (
	BET   ActionName = "PreFlop"
	RAISE ActionName = "Flop"
	CHECK ActionName = "Turn"
	FOLD  ActionName = "River"
)

type Model interface{}

type ActionModel struct {
	Name   string
	Amount float64
}

type CardModel struct {
	Suit   string
	Number string
}

type PlayerStatusModel struct {
	Cards []CardModel
	Stack float64
	InPot float64
}

type HandStatusModel struct {
	ButtonPlayer int
	CurrPlayer   int
	Stacks       []float64
	Players      []PlayerStatusModel

	CurrPhase  string
	LastAction ActionModel
}

type HandModel struct {
	StartingStatus HandStatusModel
	Actions        []ActionModel
}

type GameModel struct {
	Players       []string
	SB            int
	StartingChips int
	BlindsTimer   int
	BuyInt        int

	Hands []HandModel
	Id    string
}
