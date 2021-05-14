package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readFile(output interface{}) error {
	f, err := os.Open("list_db.txt")
	if err != nil {
		return err
	}

	if err := json.NewDecoder(f).Decode(&output); err != nil {
		return fmt.Errorf("failed to json unmarshal the response: %w", err)
	}

	return nil
}
