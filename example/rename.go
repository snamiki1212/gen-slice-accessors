package main

//go:generate go run ../. --entity Account --slice Accounts --input rename.go --output rename_gen.go --rename=AccountID:GetAccountIDs --rename=Age:AgeList
type Account struct {
	AccountID string
	Age       int64
}

type Accounts []Account
