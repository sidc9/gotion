package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kr/pretty"
)

type Database struct {
	CreatedTime    string     `json:"created_time"`
	ID             string     `json:"id"`
	LastEditedTime string     `json:"last_edited_time"`
	Object         string     `json:"object"`
	Properties     Properties `json:"properties"`
}

func (*Database) TypeName() string {
	return ObjectTypeDatabase
}

func (d *Database) String() string {
	return ""
}

func ListDatabases() {
	type ListResponse struct {
		Response
		Results []*Database `json:"results"`
	}

	var resp ListResponse

	// if err := makeRequest(http.MethodGet, "databases", nil, &resp); err != nil {
	if err := readFile(&resp); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// fmt.Println(resp)
	pretty.Println(resp)
}

type Property struct {
	ID    string      `json:"id"`
	Type  string      `json:"type"`
	Name  string      `json:"-"`
	Value interface{} `json:-`
}

type Properties []*Property

func (p Properties) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	for k, v := range m {
		prop := &Property{Name: k}
		for kk, vv := range v.(map[string]interface{}) {
			switch kk {
			case "id":
				prop.ID = vv.(string)
			case "type":
				prop.Type = vv.(string)
			default:
				prop.Value = vv
			}
		}

		p = append(p, prop)
	}

	return nil
}
