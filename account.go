package main

import (
)

// Golang technique here is encoding.
type Account struct {
    AccountID   int64  `spanner:"AccountID" csv:"AccountID"`
}
