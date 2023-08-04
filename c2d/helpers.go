package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/archsh/go.xql"
	log "github.com/sirupsen/logrus"
	//"github.com/archsh/go.xql/dialects/postgres"
)

type JSONDictionary map[string]interface{}
type JSONDictionaryArray []JSONDictionary

//type HSTOREDictionary postgres.HSTORE
//type StringArray postgres.StringArray
//type IntegerArray postgres.IntegerArray
//type SmallIntegerArray postgres.SmallIntegerArray
//type BoolArray postgres.BoolArray

func (p *JSONDictionary) Scan(src interface{}) error {
	if nil == src {
		*p = nil
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	var i JSONDictionary
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	*p = i
	return nil
}

func (p JSONDictionary) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	log.Debugln("JSONDictionary.Value:>", j, ":::", err)
	return j, err
}

func (p JSONDictionary) Declare(props xql.PropertySet) string {
	return "jsonb"
}

func (p *JSONDictionaryArray) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	var i JSONDictionaryArray
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	*p = i
	return nil
}

func (p JSONDictionaryArray) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	log.Debugln("JSONDictionaryArray.Value:>", j, ":::", err)
	return j, err
}

func (p JSONDictionaryArray) Declare(props xql.PropertySet) string {
	return "jsonb"
}
