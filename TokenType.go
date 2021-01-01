package main

type finalState int

//TokenType is
type TokenType struct {
	finalState int
}

//NewTokenType is
func NewTokenType(finalState int) *TokenType { //TokenType's constructor
	p := new(TokenType)
	p.finalState = finalState
	return p
}
