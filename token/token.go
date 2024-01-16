package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1234567890

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG = "!"
	ASTRISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string] TokenType {
	"fn":FUNCTION,
	"let":LET,
}

func LookupIdent(ident string) TokenType {
	if value, ok := keywords[ident]; ok{
		return value
	}
	return IDENT
}