package expr

const TypeValue = "Value"

type value[T any] struct {
	value T
}

// Value builds an expression that returns a constant value.
//
// Reflection of the built expression returns [TypeValue] and the constant value.
func Value[T any](v T) Expr[T] {
	return &value[T]{value: v}
}

func (e *value[T]) Eval(args ...any) (T, error) {
	return e.value, nil
}

func (e *value[T]) Reflect() (string, any) {
	return TypeValue, e.value
}
