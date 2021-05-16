package main

type Sort struct {
	Property  string    `json:"property,omitempty"`
	Timestamp string    `json:"timestamp,omitempty"`
	Direction Direction `json:"direction,omitempty"`
}

type Direction string

const (
	SortAscending  Direction = "ascending"
	SortDescending Direction = "descending"
)

func NewPropertySort(property string, direction Direction) *Sort {
	return &Sort{
		Property:  property,
		Direction: direction,
	}
}

func NewTimestampSort(timestamp string, direction Direction) *Sort {
	return &Sort{
		Timestamp: timestamp,
		Direction: direction,
	}
}
