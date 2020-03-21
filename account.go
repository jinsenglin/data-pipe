package main

import (
    "github.com/icrowley/fake"
)

// Golang technique here is encoding and marshalling.
type Account struct {
    AccountID   int64  `spanner:"AccountID" csv:"AccountID"`
    Name        string `spanner:"Name" csv:"Name"`
}

func NewAccount(id int64) *Account {
	return &Account{
		AccountID:   id,
		Name:        fake.FullName(),
	}
}
