package money

// Expression money's expression interface
type Expression interface {
	Plus(Expression) Expression
	Reduce(*Bank, string) *Money
	Times(int) Expression
}
