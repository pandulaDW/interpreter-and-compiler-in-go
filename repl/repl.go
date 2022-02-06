package repl

import (
	"bufio"
	"fmt"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/lexer"
	"github.com/pandulaDW/interpreter-and-compiler-in-go/token"
	"io"
)

const PROMPT = ">> "

// Start starts a repl and waits for user input.
//
// It tokenizes the input and returns to the given output writer
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				panic(err)
			}
		}
	}
}
