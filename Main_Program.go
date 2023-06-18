package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	IF TokenType = iota
	ELSE
	VAR
	IDENTIFIER
	INTEGER
	OPERATOR
)

type Token struct {
	Type  TokenType
	Value string
}

func lexer(input string) []Token {
	var tokens []Token

	var currentToken strings.Builder
	for _, char := range input {
		if unicode.IsSpace(char) {
			if currentToken.Len() > 0 {
				tokens = append(tokens, identifyToken(currentToken.String()))
				currentToken.Reset()
			}
			continue
		}

		if char == '=' || char == '+' || char == '-' || char == '*' || char == '/' {
			if currentToken.Len() > 0 {
				tokens = append(tokens, identifyToken(currentToken.String()))
				currentToken.Reset()
			}
			tokens = append(tokens, Token{Type: OPERATOR, Value: string(char)})
			continue
		}

		currentToken.WriteRune(char)
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, identifyToken(currentToken.String()))
	}

	return tokens
}

func identifyToken(token string) Token {
	switch token {
	case "if":
		return Token{Type: IF, Value: token}
	case "else":
		return Token{Type: ELSE, Value: token}
	case "var":
		return Token{Type: VAR, Value: token}
	default:
		if _, err := strconv.Atoi(token); err == nil {
			return Token{Type: INTEGER, Value: token}
		} else {
			return Token{Type: IDENTIFIER, Value: token}
		}
	}
}

func parser(tokens []Token) string {
	var generatedCode strings.Builder

	i := 0
	for i < len(tokens) {
		token := tokens[i]

		switch token.Type {
		case IF:
			generatedCode.WriteString("if ")
			i++
			for i < len(tokens) && tokens[i].Type != ELSE {
				generatedCode.WriteString(tokens[i].Value + " ")
				i++
			}
			generatedCode.WriteString("{\n")
			if i < len(tokens) && tokens[i].Type == ELSE {
				generatedCode.WriteString("} else {\n")
				i++
			} else {
				generatedCode.WriteString("}\n")
			}
		case VAR:
			generatedCode.WriteString("var ")
			i++
			for i < len(tokens) && tokens[i].Type != IF && tokens[i].Type != ELSE {
				generatedCode.WriteString(tokens[i].Value + " ")
				i++
			}
			generatedCode.WriteString("\n")
		default:
			generatedCode.WriteString(token.Value + " ")
			i++
			for i < len(tokens) && tokens[i].Type != IF && tokens[i].Type != ELSE {
				generatedCode.WriteString(tokens[i].Value + " ")
				i++
			}
			generatedCode.WriteString("\n")
		}
	}

	return generatedCode.String()
}

func main() {
	fmt.Println("*******************************************************************************************************************************************")
	input := `var x = 5 if x > 3 { var y = 10 } else { var z = 20 }`
	tokens := lexer(input)
	fmt.Println("Berikut adalah hasil dari fungsi lexer yang memecah input string menjadi token-token yang dapat diidentifikasi oleh parser.")
	fmt.Println(tokens)
	fmt.Println("*******************************************************************************************************************************************")
	generatedCode := parser(tokens)
	fmt.Println(generatedCode)
}
