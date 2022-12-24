package repl

import (
	"BubblePL/lexer"
	"BubblePL/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "🫧>> "

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
		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tk)
		}
	}
}
