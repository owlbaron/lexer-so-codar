package tokens

type TokenType string

const (
  IDENTIFIER TokenType = "IDENTIFIER"
  LITERAL TokenType = "LITERAL"
  OPERATOR TokenType = "OPERATOR"
  GROUPER TokenType = "GROUPER"
  WS TokenType = "WS"
  NL TokenType = "NL"
  UNKNOWN TokenType = "UNKNOWN"
)

type Token struct {
  Type   TokenType
  Value  string
  Line   uint
  Column uint
}
