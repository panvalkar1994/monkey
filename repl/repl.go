package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/panvalkar1994/monkey/evaluator"
	"github.com/panvalkar1994/monkey/lexer"
	"github.com/panvalkar1994/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		lex := lexer.New(line)
		p := parser.New(lex)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
