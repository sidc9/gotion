package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadFile(output interface{}, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(f).Decode(&output); err != nil {
		return fmt.Errorf("failed to json unmarshal the response: %w", err)
	}

	return nil
}
