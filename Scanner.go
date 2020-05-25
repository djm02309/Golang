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

//Scanner is
type Scanner struct {
	transM [4][128]int
	source string
}

//NewScanner is
func NewScanner(source string) *Scanner {
	p := new(Scanner)
	// tm := [4][128]int
	p.transM = initTM()
	p.source = source
	return p
}

func initTM() [4][128]int {
	var tm [4][128]int
	for i := 0; i < 4; i++ {
		for j := 0; i < 128; j++ {
			tm[i][j] = -1
		}
	}
	// digit = 2, '-' = 1, alpha = 3
	for i := '0'; i <= '9'; i++ {
		tm[0][i] = 2
	}
	tm[0]['-'] = 1

	for i := 'a'; i <= 'z'; i++ {
		tm[0][i] = 3
	}
	for i := 'A'; i <= 'Z'; i++ {
		tm[0][i] = 3
	}
	// '-' -> digit = 2
	for i := '0'; i <= '9'; i++ {
		tm[1][i] = 2
	}
	// digit -> digit = 2
	for i := '0'; i <= '9'; i++ {
		tm[2][i] = 2
	}
	// alpha -> digit, alpha = 3
	for i := '0'; i <= '9'; i++ {
		tm[3][i] = 3
	}
	for i := 'a'; i <= 'z'; i++ {
		tm[3][i] = 3
	}

	for i := 'A'; i <= 'Z'; i++ {
		tm[3][i] = 3
	}

	return tm
}
func nextToken() Token {
	stateOld := 0
	stateNew

	// if(){

	// }

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
