package expr

import (
	"fmt"
	"reflect"
)

const TypeRef = "Ref"

type ref[T any] struct {
	argIndex  int
	argType   reflect.Type
	fieldPath []string
}

// Ref builds an expression that returns the value of an argument or field of an argument.
//
// The target must be a [Build] argument or a pointer to a field of an [Build] argument.
//
// Referencing nested field is supported. However, it requires that the inner struct fields must not be pointers.
//
// Reflection of the built expression returns [TypeRef] and a [3]any array containing the argument index, argument
// type, and field path. The argument index is a zero based int indicating which argument it references. The argument
// type is an indirected [reflect.Type] of the referenced argument. The field path is a []string. When an argument is
// directly referenced, the field path is empty. When a field or nested field is referenced, the field path contains
// the names of the fields in order.
func Ref[T any](target *T) Expr[T] {
	argIndex, argType, fieldPath := find(target)
	if argIndex < 0 {
		panic("invalid reference")
	}
	return &ref[T]{argIndex, argType, fieldPath}
}

func (e *ref[T]) Eval(args ...any) (T, error) {
	if len(args) <= e.argIndex {
		var t T
		return t, ErrTooFewArguments
	}

	v := reflect.ValueOf(args[e.argIndex])
	for v.Type() != e.argType && v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Type() != e.argType {
		var t T
		return t, fmt.Errorf("%w at index %d", ErrWrongArgumentType, e.argIndex)
	}

	for i := 0; i < len(e.fieldPath); i++ {
		v = v.FieldByName(e.fieldPath[i])
	}

	return v.Interface().(T), nil
}

func (e *ref[T]) Reflect() (string, any) {
	return TypeRef, [3]any{e.argIndex, e.argType, e.fieldPath}
}
