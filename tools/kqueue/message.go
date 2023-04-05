package kqueue

const (
	RelUpdateTableGroup = "group"
	RelUpdateTableRole  = "role"
)

type RelUpdateMqMsg struct {
	Table string `json:"table"`
	Id    int64  `json:"id"`
}
