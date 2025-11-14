package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joevtap/monkey/internal/lexer"
	"github.com/joevtap/monkey/internal/token"
)

func main() {
	fmt.Println(`This is the Monkey programming language!
Feel free to type in commands`)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(">> ")
	for scanner.Scan() {
		input := scanner.Text()

		l := lexer.New(input)

		for {
			tok := l.NextToken()

			if tok.Type == token.EOF {
				break
			}

			fmt.Printf("%+v\n", tok)
		}

		fmt.Printf(">> ")
	}
}
