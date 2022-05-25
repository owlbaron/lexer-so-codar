package literal

import (
  . "main/automata"
  "unicode"
)

func NewLiteralAutomata() *Automata {
  aut := NewAutomata(0, false)

  aut.AddState(1, true)
  aut.AddState(2, false)
  aut.AddState(3, true)
  aut.AddState(4, false)
  aut.AddState(5, false)
  aut.AddState(6, true)

  aut.AddTransition(0, func (input rune) bool { return unicode.IsNumber(input) }, 1)
  aut.AddTransition(0, func (input rune) bool { return input == '.' }, 2)
  aut.AddTransition(1, func (input rune) bool { return unicode.IsNumber(input) }, 1)
  aut.AddTransition(1, func (input rune) bool { return input == '.' }, 2)
  aut.AddTransition(1, func (input rune) bool { return input == 'e' }, 4)
  aut.AddTransition(2, func (input rune) bool { return unicode.IsNumber(input) }, 3)
  aut.AddTransition(3, func (input rune) bool { return unicode.IsNumber(input) }, 3)
  aut.AddTransition(3, func (input rune) bool { return input == 'e' }, 4)
  aut.AddTransition(4, func (input rune) bool { return input == '-' || input == '+' }, 5)
  aut.AddTransition(4, func (input rune) bool { return unicode.IsNumber(input) }, 6)
  aut.AddTransition(5, func (input rune) bool { return unicode.IsNumber(input) }, 6)
  aut.AddTransition(6, func (input rune) bool { return unicode.IsNumber(input) }, 6)
  
  return aut
}