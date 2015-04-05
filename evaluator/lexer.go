package evaluator

import "fmt"

type ExprLexer struct {
}

type lexerInput struct {
	text string
	pos  int
}

func NewLexerInput(text string, pos int) *lexerInput {
	return &lexerInput{
		text: text,
		pos:  pos,
	}
}

func (input *lexerInput) Current() (uint8, bool) {
	if input.pos >= len(input.text) {
		return 0, false
	}
	return input.text[input.pos], true
}

func (input *lexerInput) Consume() bool {
	if input.pos < len(input.text) {
		input.pos += 1
		return true
	}
	return false
}

func (input *lexerInput) ConsumeChar(char uint8) bool {
	if input.pos < len(input.text) && input.text[input.pos] == char {
		input.pos += 1
		return true
	}
	return false
}

func (input *lexerInput) End() bool {
	return input.pos >= len(input.text)
}

func isLetter(char uint8) bool {
	return char >= 'a' && char <= 'z' || (char >= 'A' && char <= 'Z')
}

func isDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}

func isWhitespace(char uint8) bool {
	return char == ' ' || char == '\n' || char == '\r'
}

func (lexer *ExprLexer) lexExpression(expression string) []*ExprToken {
	input := NewLexerInput(expression, 0)
	tokens := make([]*ExprToken, 0)

	var currentPos int
	var currentChar uint8
	var ok bool

	for !input.End() {
		currentPos = input.pos
		currentChar, ok = input.Current()
		if !ok {
			panic("Could not get current character")
		}

		if isDigit(currentChar) {
			tokens = append(tokens, lexer.lexNumber(input))
		} else if isLetter(currentChar) {
			tokens = append(tokens, lexer.lexState(input))
		} else if isWhitespace(currentChar) {
			input.Consume()
		} else {
			tokens = append(tokens, lexer.lexSymbol(input))
		}

		// We must consume some characters every iteration
		if input.pos <= currentPos {
			panic("No characters consumed during iteration")
		}
	}
	return tokens
}

func (*ExprLexer) lexNumber(input *lexerInput) *ExprToken {

	startPos := input.pos
	var ok bool
	var char uint8
	for char, ok = input.Current(); isDigit(char) && ok; {
		input.Consume()
	}

	return NewExprToken(numberType, input.text[startPos:input.pos])
}

func (*ExprLexer) lexState(input *lexerInput) *ExprToken {
	startPos := input.pos
	var ok bool
	var char uint8
	for char, ok = input.Current(); isLetter(char) && ok; {
		input.Consume()
	}
	return NewExprToken(stateType, input.text[startPos:input.pos])
}

func (*ExprLexer) lexSymbol(input *lexerInput) *ExprToken {
	char, ok := input.Current()
	if !ok {
		panic("Could not get current character")
	}

	if char == '+' || char == '-' || char == '*' || char == '/' || char == '^' {
		input.Consume()
		return NewExprToken(operatorType, fmt.Sprintf("%d", char))
	} else if char == '(' {
		input.Consume()
		return NewExprToken(lParenType, "(")
	} else if char == ')' {
		return NewExprToken(rParenType, ")")
	}

	if input.End() {
		panic("Unexpected end of expression")
	}
	currentChar, _ := input.Current()
	panic(fmt.Sprintf("Unexpected character %s at position %d", currentChar, input.pos))

}
