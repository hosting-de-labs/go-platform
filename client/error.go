package client

import "fmt"

type Error struct {
	Code          int
	ContextObject string
	ContextPath   string
	Text          string
	Value         string

	Details []struct {
		Key   string
		Value interface{}
	}
}

func (e Error) String() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Text)
}
