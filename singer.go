package main

import (
    "github.com/icrowley/fake"
)

// Golang technique here is encoding and marshalling.
type Signer struct {
    SignerID    string  `spanner:"SignerID" csv:"SignerID"`
    Name        string `spanner:"Name" csv:"Name"`
}

func NewSigner(id string) *Signer {
	return &Signer{
		SignerID:    id,
		Name:        fake.FullName(),
	}
}
