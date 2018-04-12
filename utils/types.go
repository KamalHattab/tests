package utils

import (
	"fmt"
	"time"
)

//Inventory struct
type Inventory struct {
	Item   string                 `json:"item" bson:"item"`
	Qty    uint16                 `json:"qty" bson:"qty"`
	Size   map[string]interface{} `json:"size" bson:"size"`
	Status string                 `json:"status" bson:"status"`
}

//Game Struct
type Game struct {
	Winner       string    `bson:"winner"`
	OfficialGame bool      `bson:"official_game"`
	Location     string    `bson:"location"`
	StartTime    time.Time `bson:"start"`
	EndTime      time.Time `bson:"end"`
	Players      []Player  `bson:"players"`
}

//Player Struct
type Player struct {
	Name   string    `bson:"name"`
	Decks  [2]string `bson:"decks"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}

// String function applied on Game struct
func (g Game) String() string {

	return fmt.Sprint("{Winner:", g.Winner, ",", "OfficialName:", g.OfficialGame, ",", "Location:", g.Location, ",", "StartTime:", g.StartTime, ",",
		"EndTime:", g.EndTime, ",", "Players:", g.Players, "}")

}

//Result struct
type Result struct {
	OfficialGame bool `bson:"official_game"`
}

//const
const (
	Username            = ""
	Password            = ""
	Database            = "db_test"
	Collection          = "table_test"
	CollectionInventory = "inventory_go"
)

