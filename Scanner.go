package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/scanner"
)

//Scanner is
type Scanner struct {
	//Golang 에서는 배열을 꼭 크기 지정 필요 크기지정없으면 slice됨
	transM [4][128]int
	source string
}

var st scanner.Scanner

func (s *st) NextToken(){
	return
}

//NewScanner is constructor
func NewScanner(source string) *Scanner {
	p := new(Scanner)
	// tm := [4][128]int
	p.transM = initTM()
	if source == nil {
		p.source=""
	}else{
		p.source=source
	}
	st = new
	return p
}
func (s *Scanner) tokenize() []Token {
	var list []Token //slice 선언(java의 list 역할할것)
	for {            //조건식 추가 필요
		list = append(list, nextToken())
	}
	return list
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
	stateNew :=0

	if(!){

	}

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
