package main

import (
	"encoding/json"
)

func MovesService() ([]byte, error) {
	daoResponse := MovesDao()
	res, err := json.Marshal(daoResponse)
	return res, err
}
