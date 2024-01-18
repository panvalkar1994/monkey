package parser

import (
	"fmt"

	"github.com/panvalkar1994/monkey/ast"
	"github.com/panvalkar1994/monkey/lexer"
	"github.com/panvalkar1994/monkey/token"
)

type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
	errors []string
}


func New(l *lexer.Lexer) *Parser {
	p := &Parser{l:l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p* Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",t,  p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	st := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT){
		return nil
	}
	st.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping till we get semicolon
	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return st
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	st := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return st
	
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		st := p.parseStatement()
		if st!=nil {
			prog.Statements = append(prog.Statements, st)
		}
		p.nextToken()
	}
	return prog
}