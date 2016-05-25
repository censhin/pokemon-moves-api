package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Moves struct{
	Moves	[]Move `json:"moves"`
}

type Move struct {
	Name       string   `json:"name"`
	Number     int      `json:"number"`
	Type       string   `json:"type"`
	Category   string   `json:"category"`
	Contest    string   `json:"contest"`
	Pp         int      `json:"pp"`
	Power      int      `json:"power"`
	Accuracy   int      `json:"accuracy"`
	Generation string   `json:"generation"`
	Tm         []string `json:"tm"`
	Hm         []string `json:"hm"`
	Tutor      bool     `json:"tutor"`
}

func MovesDao() Moves {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("pokedex").C("moves")

	result := []Move{}
	iter := collection.Find(nil).Iter()
	err = iter.All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return Moves{result}
}

func MoveDao(name string) Move {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("pokedex").C("moves")

	result := Move{}
	err = collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
