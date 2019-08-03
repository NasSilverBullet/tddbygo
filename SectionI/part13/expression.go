package money

// Expression money's expression
type Expression interface {
	Plus(*Money) Expression
}
