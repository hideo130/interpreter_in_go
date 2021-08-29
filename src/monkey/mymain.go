package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"monkey/lexer"
)
func mymain(){

	fmt.Println("ファイル読み取り処理を開始します")
    // ファイルをOpenする
    f, err := os.Open("generate_test.txt")
    if err != nil{
        fmt.Println("error")
    }
    defer f.Close()
    // 一気に全部読み取り
    b, err := ioutil.ReadAll(f)
	if err != nil{
		fmt.Print(err)
	}
	input := string(b)
	literals := strings.Split(input, "\n")
	l := lexer.New(input)
	for range literals {
		tok := l.NextToken()
		fmt.Printf("{token.%s,\"%s\"},\n", tok.Type, tok.Literal)
		
	}
}