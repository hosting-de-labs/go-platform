package client

type Error struct {
	Code          int
	ContextObject string
	ContextPath   string
	Details       map[string]string
	Text          string
	Value         string
}
