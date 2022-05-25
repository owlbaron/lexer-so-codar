package lexer

import (
  "strings"
  
  "main/automata/tokens/grouper"
  "main/automata/tokens/identifier"
  "main/automata/tokens/literal"
  "main/automata/tokens/operator"
  "main/automata/tokens/ws"
  "main/automata/tokens/nl"

  . "main/automata"
  . "main/tokens"
)

type result struct {
  Accepted bool
  Type TokenType
  Value string
}

type testInput struct {
  Automata *Automata
  Type TokenType
}

type lexer struct {
  code string
  ptr uint

  line uint
  column uint
}

func NewLexer(code string) *lexer {
  return &lexer{
    code: code,
    ptr: 0,

    line: 1,
    column: 1,
  }
}

func (l *lexer) Next() *Token {
  resultChannel := make(chan result)

  tests := []testInput{
    {Automata: grouper.NewGrouperAutomata(), Type: GROUPER},
    {Automata: identifier.NewIdentifierAutomata(), Type: IDENTIFIER},
    {Automata: literal.NewLiteralAutomata(), Type: LITERAL},
    {Automata: operator.NewOperatorAutomata(), Type: OPERATOR},
    {Automata: ws.NewWSAutomata(), Type: WS},
    {Automata: nl.NewNLAutomata(), Type: NL},
  }

  for _, test := range tests {
    go l.TestAutomata(resultChannel, test.Automata, test.Type)
  }

  for i := 0; i < len(tests); i++ {
    result := <- resultChannel

    if result.Accepted {
      token := &Token{
        Type: result.Type,
        Value: result.Value,
        Line: l.line,
        Column: l.column,
      }
      
      l.ptr += uint(len(result.Value))
      if (result.Type == NL) {
        l.line++
        l.column = 1
      } else {
        l.column += uint(len(result.Value))
      }
      
      return token
    }
  }

  if (l.ptr < uint(len(l.code) - 1)) {
    token := &Token{
      Type: UNKNOWN,
      Value: string(l.code[l.ptr]),
      Line: l.line,
      Column: l.column,
    }

    l.column++
    l.ptr++

    return token
  }

  return nil
}

func (l *lexer) TestAutomata(channel chan result, automata *Automata, tokenType TokenType) {
  var sb strings.Builder
  ptr := int(l.ptr)

  for ptr < len(l.code) {
    char := rune(l.code[ptr])

    err := automata.Input(char)

    if err != nil {
      break
    }

    sb.WriteRune(char)
    ptr++
  }

  channel <- result{
    Accepted: automata.Accepted(),
    Type: tokenType,
    Value: sb.String(),
  }
}
