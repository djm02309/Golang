package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func main() {
	file, err := os.Open("/path/to/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello world")
}
