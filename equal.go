package expr

const TypeEqual = "Equal"

type equal[T comparable] struct {
	a Expr[T]
	b Expr[T]
}

// Equal builds an expression that compares two inner expressions for equality.
//
// Reflection of the built expression returns [TypeEqual] and a [2]Expr[T] array containing the two inner expressions.
func Equal[T comparable](a Expr[T], b Expr[T]) Expr[bool] {
	return &equal[T]{a, b}
}

// EqualRR is a sugar function for Equal, where both arguments are references.
func EqualRR[T comparable](a *T, b *T) Expr[bool] {
	return Equal(Ref(a), Ref(b))
}

// EqualRV is a sugar function for Equal, where the first argument is a reference and the second is a value.
func EqualRV[T comparable](r *T, v T) Expr[bool] {
	return Equal(Ref(r), Value(v))
}

func (e *equal[T]) Eval(args ...any) (bool, error) {
	va, err := e.a.Eval(args...)
	if err != nil {
		return false, err
	}
	vb, err := e.b.Eval(args...)
	if err != nil {
		return false, err
	}
	return va == vb, nil
}

func (e *equal[T]) Reflect() (string, any) {
	return TypeEqual, [2]Expr[T]{e.a, e.b}
}
