package main

//Token is
type Token struct {
	tokentype TokenType
	lexme     string
}

const (
	//INT is
	INT finalState = 2 + iota
	//ID is
	ID
)
