package evaluator

var operators = []string{"+", "-", "*", "/", "^"}

type ExpressionParser struct {
}

type parserInput struct {
	tokens []*ExprToken
	pos    int
}

func (p *parserInput) Next() *ExprToken {
	if p.pos < len(p.tokens) {
		p.pos += 1
		return p.tokens[p.pos-1]
	}
	return nil
}

func (p *parserInput) StepBack() {
	if p.pos > 0 {
		p.pos -= 1
	}
}

func (p *parserInput) End() bool {
	return p.pos >= len(p.tokens)
}

func NewParserInput(tokens []*ExprToken, pos int) *parserInput {
	return &parserInput{
		tokens: tokens,
		pos:    pos,
	}
}

type parseFunc func(input *parserInput) *ExpressionNode

func (parser *ExpressionParser) parseExpression(input *parserInput) *ExpressionNode {
	return parser.parseBinaryOperator(input)
}

func (parser *ExpressionParser) parseBinaryOperator(input *parserInput, parseSubtree parseFunc, operators []string) *ExpressionNode {
	lhs := parseSubtree(input)

	operatorToken := input.next()

	if operatorToken == nil {
		return lhs
	}
	if operatorToken.tokenType != operatorType || !operators.contains(operatorToken.text) {
		input.StepBack()
		return lhs
	}

	var rhs, result *ExpressionNode
	for true {
		rhs = parseSubtree(input)

		// We have successfully parsed an operator expression
		result = NewOperatorExpressionNode(lhs, operatorToken.text, rhs)

		operatorToken = input.next()
		if operatorToken == nil {
			break
		}
		if operatorToken.tokenType != operatorType || !operators.contains(operatorToken.text) {
			input.StepBack()
			break
		}

		// we have another operator expression for which to parse the rhs
		lhs = result

	}
	return result
}

func parseTerm(input *parserInput) ExpressionNode {

	token := input.next()
	if token == nil {
		panic("Unexpected end of expression")
	}

	switch token.tokenType {
	case numberType:
		return NewNumberExpressionNode(token.text)
	case lParenType:
		subNode := parseExpression(input)
		rParenToken := input.next()

		if rParenToken == nil || rParenToken.tokenType != rParenType {
			panic("Missing closing parenthesis")
		}
		return subNode
	}
	panic("Missing term in expression")
}
