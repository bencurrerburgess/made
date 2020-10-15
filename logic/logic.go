package logic

import (
	"encoding/json"
	"io/ioutil"
)


type Example struct {
	name string
	value int
}


func LoadJsonFile(filepath string, structure interface{})  error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, structure)
	if err != nil {
		return err
	}

	return nil
}
