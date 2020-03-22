package main

import (
    "github.com/icrowley/fake"
)

// Golang technique here is struct tag for encoding and marshalling.
type Singer struct {
    SingerID    string  `spanner:"SingerID" csv:"SingerID"`
    Name        string `spanner:"Name" csv:"Name"`
}

func NewSinger(id string) *Singer {
	return &Singer{
		SingerID:    id,
		Name:        fake.FullName(),
	}
}
