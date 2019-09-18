package models

// Rectangle is used to hold the rect parameters
type Rectangle struct {
	Length  int64 `json:"length,omitempty"`
	Breadth int64 `json:"breadth,omitempty"`
}
