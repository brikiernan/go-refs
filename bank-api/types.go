package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Number    int64  `json:"account"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10_000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(10_000_000)),
	}
}
