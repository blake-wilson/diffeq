package evaluator

type ExprToken struct {
	tokenType, text string
}

const (
	numberType   = "Number"
	operatorType = "operator"
	lParenType   = "lparen"
	rParenType   = "rparen"
)

func (token *ExprToken) Equals(other *ExprToken) bool {
	return token.tokenType == other.tokenType &&
		token.text == other.text
}

func NewExprToken(tokenType, text string) *ExprToken {
	return &ExprToken{
		tokenType: tokenType,
		text:      text,
	}
}
