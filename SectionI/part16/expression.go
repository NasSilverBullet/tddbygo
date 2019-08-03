package money

// Expression money's expression
type Expression interface {
	Plus(Expression) Expression
	Reduce(*Bank, string) *Money
}
