package expr

// Expr[T] represents an expression with returning type T.
//
// Eval evaluates the expression with the given arguments and returns the result.
//
// Reflect returns the type of the expression and data associated with it.
type Expr[T any] interface {
	Eval(args ...any) (T, error)
	Reflect() (string, any)
}
