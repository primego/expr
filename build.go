package expr

import (
	"fmt"
	"reflect"
)

// Build[T] builds an Expr[T] using the provided builder function. The builder function can take any number of
// arguments and must return an Expr[T]. The arguments must be pointers and can be referenced by using [Ref]. The
// result of the builder function will be returned as the result of Build[T].
func Build[T any](b any) Expr[T] {
	rb := reflect.ValueOf(b)
	if rb.Kind() != reflect.Func {
		panic("builder must be a function")
	}

	rbt := rb.Type()
	eot := reflect.TypeOf((*Expr[T])(nil)).Elem()
	if rbt.NumOut() != 1 || rbt.Out(0) != eot {
		panic(fmt.Sprintf("builder must return an expr.%s", eot.Name()))
	}

	ctx := &context{}
	contexts.add(ctx)
	defer contexts.delete(ctx)

	args := make([]reflect.Value, rbt.NumIn())
	for i := 0; i < rbt.NumIn(); i++ {
		argType := rbt.In(i)
		if argType.Kind() != reflect.Ptr {
			panic("builder arguments must be pointers")
		}
		arg := reflect.New(argType.Elem())
		ctx.add(arg.Interface())
		args[i] = arg
	}
	results := rb.Call(args)
	return results[0].Interface().(Expr[T])
}
