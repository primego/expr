package expr_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/primego/expr"
)

func TestAlias(t *testing.T) {
	if reflect.TypeOf((*expr.Bool)(nil)) != reflect.TypeOf((*expr.Expr[bool])(nil)) {
		t.Fatal("expr.Bool is not an alias for expr.Expr[bool]")
	}
	if reflect.TypeOf((*expr.Float32)(nil)) != reflect.TypeOf((*expr.Expr[float32])(nil)) {
		t.Fatal("expr.Float32 is not an alias for expr.Expr[float32]")
	}
	if reflect.TypeOf((*expr.Float64)(nil)) != reflect.TypeOf((*expr.Expr[float64])(nil)) {
		t.Fatal("expr.Float64 is not an alias for expr.Expr[float64]")
	}
	if reflect.TypeOf((*expr.Int)(nil)) != reflect.TypeOf((*expr.Expr[int])(nil)) {
		t.Fatal("expr.Int is not an alias for expr.Expr[int]")
	}
	if reflect.TypeOf((*expr.Int8)(nil)) != reflect.TypeOf((*expr.Expr[int8])(nil)) {
		t.Fatal("expr.Int8 is not an alias for expr.Expr[int8]")
	}
	if reflect.TypeOf((*expr.Int16)(nil)) != reflect.TypeOf((*expr.Expr[int16])(nil)) {
		t.Fatal("expr.Int16 is not an alias for expr.Expr[int16]")
	}
	if reflect.TypeOf((*expr.Int32)(nil)) != reflect.TypeOf((*expr.Expr[int32])(nil)) {
		t.Fatal("expr.Int32 is not an alias for expr.Expr[int32]")
	}
	if reflect.TypeOf((*expr.Int64)(nil)) != reflect.TypeOf((*expr.Expr[int64])(nil)) {
		t.Fatal("expr.Int64 is not an alias for expr.Expr[int64]")
	}
	if reflect.TypeOf((*expr.String)(nil)) != reflect.TypeOf((*expr.Expr[string])(nil)) {
		t.Fatal("expr.String is not an alias for expr.Expr[string]")
	}
	if reflect.TypeOf((*expr.Time)(nil)) != reflect.TypeOf((*expr.Expr[time.Time])(nil)) {
		t.Fatal("expr.Time is not an alias for expr.Expr[time.Time]")
	}
	if reflect.TypeOf((*expr.Uint)(nil)) != reflect.TypeOf((*expr.Expr[uint])(nil)) {
		t.Fatal("expr.Uint is not an alias for expr.Expr[uint]")
	}
	if reflect.TypeOf((*expr.Uint8)(nil)) != reflect.TypeOf((*expr.Expr[uint8])(nil)) {
		t.Fatal("expr.Uint8 is not an alias for expr.Expr[uint8]")
	}
	if reflect.TypeOf((*expr.Uint16)(nil)) != reflect.TypeOf((*expr.Expr[uint16])(nil)) {
		t.Fatal("expr.Uint16 is not an alias for expr.Expr[uint16]")
	}
	if reflect.TypeOf((*expr.Uint32)(nil)) != reflect.TypeOf((*expr.Expr[uint32])(nil)) {
		t.Fatal("expr.Uint32 is not an alias for expr.Expr[uint32]")
	}
	if reflect.TypeOf((*expr.Uint64)(nil)) != reflect.TypeOf((*expr.Expr[uint64])(nil)) {
		t.Fatal("expr.Uint64 is not an alias for expr.Expr[uint64]")
	}
}
