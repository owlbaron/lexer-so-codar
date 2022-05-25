package main

import (
	"fmt"
  "io/ioutil"
  
  "main/lexer"
)

func main() {
  var filePath string

  fmt.Println("Digite o caminho do arquivo '.sc'")
  fmt.Scanf("%s", &filePath)

  content, err := ioutil.ReadFile(filePath)
  if err != nil {
      fmt.Println(err)
  }
  
  l := lexer.NewLexer(string(content))

  token := l.Next()

  for token != nil {
    fmt.Println(token)
    
    token = l.Next()
  }
}
