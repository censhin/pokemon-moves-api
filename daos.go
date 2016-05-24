package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Move struct {
	Name string
	Number int
	Type string
	Category string
	Contest string
	Pp int
	Power int
	Accuracy int
	Generation string
	Tm []string
	Hm []string
	Tutor bool
}

func MovesDao() ([]Move) {
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

	return result
}

func MoveDao(name string) (Move) {
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
