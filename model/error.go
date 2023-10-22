package model

type Error struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
