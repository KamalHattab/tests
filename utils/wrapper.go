package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2" 
	"gopkg.in/mgo.v2/bson"
)

//NewPlayer Struct
func NewPlayer(name, firstDeck, secondDeck string, points, place uint8) Player {
	return Player{
		Name:   name,
		Decks:  [2]string{firstDeck, secondDeck},
		Points: points,
		Place:  place,
	}
}

//CreatePlayers Function
func CreatePlayers() []Player {
	return []Player{
		NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
		NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
		NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
		NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		NewPlayer("Kamal", "Aliens", "Dinosaurs", 17, 3),
		NewPlayer("Rodrigo", "Spies", "Leprechauns", 10, 4),
	}
}

var isDropMe, insertMore = true, false

//QueryDocuments function
func QueryDocuments(session *mgo.Session) {
	c := session.DB(Database).C(CollectionInventory)
	if insertMore {
		inv1 := Inventory{Item: "journal", Qty: 25, Status: "A", Size: map[string]interface{}{"h": 14, "w": 21, "uom": "cm"}}
		inv2 := Inventory{Item: "notebook", Qty: 50, Status: "A", Size: map[string]interface{}{"h": 8.5, "w": 11, "uom": "in"}}
		inv3 := Inventory{Item: "paper", Qty: 100, Status: "D", Size: map[string]interface{}{"h": 8.5, "w": 11, "uom": "in"}}
		inv4 := Inventory{Item: "planner", Qty: 75, Status: "D", Size: map[string]interface{}{"h": 22.85, "w": 30, "uom": "cm"}}
		inv5 := Inventory{Item: "postcard", Qty: 45, Status: "A", Size: map[string]interface{}{"h": 10, "w": 15.25, "uom": "cm"}}
		err := c.Insert(inv1, inv2, inv3, inv4, inv5)
		if err != nil {
			panic(err)
		}
		fmt.Println("inserted with success")
	}
	var inventories []Inventory
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< All Documents inside collections")
	c.Find(nil).All(&inventories)
	b, _ := json.Marshal(inventories)
	fmt.Println("all as Json arrary\n", string(b))

	for i := range inventories {
		b, _ = json.Marshal(inventories[i])
		fmt.Println(string(b))
	}

}

//Aggregate Function
func Aggregate(session *mgo.Session) {
	c := session.DB(Database).C(CollectionInventory)
	pipeLine := []bson.M{
		bson.M{"$match": bson.M{"qty": bson.M{"$gte": 25}}},
		//bson.M{"$project": bson.M{"status": 1}},
		bson.M{"$group": bson.M{"_id": "$status", "total": bson.M{"$sum": "$qty"}, "count": bson.M{"$sum": 1}}},
	}
	var result struct { //interface{}
		ID string `bson:"_id" json:"_id"`
		//Status string `bson:"status"`
		Total int `bson:"total"`
		Count int `bson:"count"`
		//City      string `bson:"city"`
		//Address   string `bson:"address"`
		//NoteCount int    `bson:"notecount"`
	}
	//var result interface{}
	iter := c.Pipe(pipeLine).Iter()
	defer iter.Close()
	//var inv Inventory
	for iter.Next(&result) {
		fmt.Printf("%+v\n", result)
	}
	if iter.Err() != nil {
		log.Println(iter.Err())
	}
}

//BaseExample Function
func BaseExample(session *mgo.Session) {
	game := Game{
		Winner:       "Kamal",
		OfficialGame: true,
		Location:     "Austin",
		StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
		EndTime:      time.Now(),
		Players:      CreatePlayers(),
	}

	c := session.DB(Database).C(Collection)
	if insertMore {
		//insert
		if err := c.Insert(game); err != nil {
			panic(err)
		} else {
			fmt.Println("inserted with success")
		}
	}
	// find and count
	player := "Kamal"
	gamesWon, err := c.Find(bson.M{"winner": player}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", player, gamesWon)

	// find one
	var result Game
	err = c.Find(bson.M{"winner": player, "location": "Austin"}).Select(bson.M{"official_game": 1}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)

	// find all games ordered in reverse order by the field start
	var results []Game
	err = c.Find(nil).Sort("-start").All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", len(results))

	//update
	newPlayer := "John"
	selector := bson.M{"winner": player}
	updator := bson.M{"$set": bson.M{"winner": newPlayer}}
	if err := c.Update(selector, updator); err != nil {
		panic(err)
	}

	// //remove All
	// info, err := c.RemoveAll(nil)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("removed", info.Removed)

	err = c.Find(nil).Sort("-start").All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("after updating")
	for i := range results {
		fmt.Println("results[", i, "]", results[i])
	}

}

//UpdateValueArray Function
func UpdateValueArray(session *mgo.Session) {
	type student struct {
		ID     int       `bson:"_id" json:"_id"`
		Grades []float64 `bson:"grades" json:"grades"`
	}
	c := session.DB(Database).C("students")
	//m1 := bson.M{"grades.grade": 85}
	//m1 := bson.M{"_id": 5, "grades": bson.M{"$elemMatch": bson.M{"grade": bson.M{"$lte": 90}, "mean": bson.M{"$gte": 75, "$lte": 85}}}}
	m1 := bson.M{}
	m2 := bson.M{"$inc": bson.M{"grades.$[]": -12}}

	err := c.Update(m1, m2)
	if err != nil {
		panic(err)
	}
	var result []map[string]interface{}
	c.Find(nil).All(&result)
	b, _ := json.Marshal(result)
	fmt.Print("result:", string(b))
}



