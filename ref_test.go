package expr_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/primego/expr"
)

func TestRef(t *testing.T) {
	type TestDeepNestedStruct struct {
		Float64 float64
	}

	type TestNestedStruct struct {
		Bool   bool
		Struct TestDeepNestedStruct
	}

	type TestType struct {
		Bool    bool
		Float64 float64
		Int     int
		String  string
		Time    time.Time
		Struct  TestNestedStruct
	}

	t.Run("Eval", func(t *testing.T) {
		t.Run("bool value", func(t *testing.T) {
			e := expr.Build[bool](func(arg *bool) expr.Bool {
				return expr.Ref(arg)
			})
			r, err := e.Eval(true)
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected true, got false")
			}
			r, err = e.Eval(false)
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected false, got true")
			}
		})

		t.Run("float64 value", func(t *testing.T) {
			e := expr.Build[float64](func(arg *float64) expr.Float64 {
				return expr.Ref(arg)
			})
			r, err := e.Eval(1.0)
			if err != nil {
				t.Fatal(err)
			}
			if r != 1.0 {
				t.Fatalf("expected 1.0, got %v", r)
			}
			r, err = e.Eval(2.0)
			if err != nil {
				t.Fatal(err)
			}
			if r != 2.0 {
				t.Fatalf("expected 2.0, got %v", r)
			}
		})

		t.Run("int value", func(t *testing.T) {
			e := expr.Build[int](func(arg *int) expr.Int {
				return expr.Ref(arg)
			})
			r, err := e.Eval(1)
			if err != nil {
				t.Fatal(err)
			}
			if r != 1 {
				t.Fatalf("expected 1, got %v", r)
			}
			r, err = e.Eval(2)
			if err != nil {
				t.Fatal(err)
			}
			if r != 2 {
				t.Fatalf("expected 2, got %v", r)
			}
		})

		t.Run("string value", func(t *testing.T) {
			e := expr.Build[string](func(arg *string) expr.String {
				return expr.Ref(arg)
			})
			r, err := e.Eval("foo")
			if err != nil {
				t.Fatal(err)
			}
			if r != "foo" {
				t.Fatalf(`expected "foo", got "%v"`, r)
			}
			r, err = e.Eval("bar")
			if err != nil {
				t.Fatal(err)
			}
			if r != "bar" {
				t.Fatalf(`expected "bar", got "%v"`, r)
			}
		})

		t.Run("time value", func(t *testing.T) {
			e := expr.Build[time.Time](func(arg *time.Time) expr.Time {
				return expr.Ref(arg)
			})
			now := time.Now()
			r, err := e.Eval(now)
			if err != nil {
				t.Fatal(err)
			}
			if r != now {
				t.Fatalf("expected %v, got %v", now, r)
			}
		})

		t.Run("bool field", func(t *testing.T) {
			e := expr.Build[bool](func(arg *TestType) expr.Bool {
				return expr.Ref(&arg.Bool)
			})
			r, err := e.Eval(&TestType{Bool: true})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected true, got false")
			}
			r, err = e.Eval(&TestType{Bool: false})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected false, got true")
			}
		})

		t.Run("float64 field", func(t *testing.T) {
			e := expr.Build[float64](func(arg *TestType) expr.Float64 {
				return expr.Ref(&arg.Float64)
			})
			r, err := e.Eval(&TestType{Float64: 1.0})
			if err != nil {
				t.Fatal(err)
			}
			if r != 1.0 {
				t.Fatalf("expected 1.0, got %v", r)
			}
			r, err = e.Eval(&TestType{Float64: 2.0})
			if err != nil {
				t.Fatal(err)
			}
			if r != 2.0 {
				t.Fatalf("expected 2.0, got %v", r)
			}
		})

		t.Run("int field", func(t *testing.T) {
			e := expr.Build[int](func(arg *TestType) expr.Int {
				return expr.Ref(&arg.Int)
			})
			r, err := e.Eval(&TestType{Int: 1})
			if err != nil {
				t.Fatal(err)
			}
			if r != 1 {
				t.Fatalf("expected 1, got %v", r)
			}
			r, err = e.Eval(&TestType{Int: 2})
			if err != nil {
				t.Fatal(err)
			}
			if r != 2 {
				t.Fatalf("expected 2, got %v", r)
			}
		})

		t.Run("string field", func(t *testing.T) {
			e := expr.Build[string](func(arg *TestType) expr.String {
				return expr.Ref(&arg.String)
			})
			r, err := e.Eval(&TestType{String: "foo"})
			if err != nil {
				t.Fatal(err)
			}
			if r != "foo" {
				t.Fatalf(`expected "foo", got "%v"`, r)
			}
			r, err = e.Eval(&TestType{String: "bar"})
			if err != nil {
				t.Fatal(err)
			}
			if r != "bar" {
				t.Fatalf(`expected "bar", got "%v"`, r)
			}
		})

		t.Run("time field", func(t *testing.T) {
			e := expr.Build[time.Time](func(arg *TestType) expr.Time {
				return expr.Ref(&arg.Time)
			})
			now := time.Now()
			r, err := e.Eval(&TestType{Time: now})
			if err != nil {
				t.Fatal(err)
			}
			if r != now {
				t.Fatalf("expected %v, got %v", now, r)
			}
		})

		t.Run("struct field", func(t *testing.T) {
			e := expr.Build[TestNestedStruct](func(arg *TestType) expr.Expr[TestNestedStruct] {
				return expr.Ref(&arg.Struct)
			})
			target := TestType{Struct: TestNestedStruct{Bool: true}}
			r, err := e.Eval(&target)
			if err != nil {
				t.Fatal(err)
			}
			if r != target.Struct {
				t.Fatalf("expected %v, got %v", target.Struct, r)
			}
			target = TestType{Struct: TestNestedStruct{Bool: false}}
			r, err = e.Eval(&target)
			if err != nil {
				t.Fatal(err)
			}
			if r != target.Struct {
				t.Fatalf("expected %v, got %v", target.Struct, r)
			}
		})

		t.Run("nested field", func(t *testing.T) {
			e := expr.Build[bool](func(arg *TestType) expr.Bool {
				return expr.Ref(&arg.Struct.Bool)
			})
			r, err := e.Eval(&TestType{Struct: TestNestedStruct{Bool: true}})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected true, got false")
			}
			r, err = e.Eval(&TestType{Struct: TestNestedStruct{Bool: false}})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected false, got true")
			}
		})

		t.Run("deep nested field", func(t *testing.T) {
			e := expr.Build[float64](func(arg *TestType) expr.Float64 {
				return expr.Ref(&arg.Struct.Struct.Float64)
			})
			r, err := e.Eval(&TestType{Struct: TestNestedStruct{Struct: TestDeepNestedStruct{Float64: 1.0}}})
			if err != nil {
				t.Fatal(err)
			}
			if r != 1.0 {
				t.Fatalf("expected 1.0, got %v", r)
			}
			r, err = e.Eval(&TestType{Struct: TestNestedStruct{Struct: TestDeepNestedStruct{Float64: 2.0}}})
			if err != nil {
				t.Fatal(err)
			}
			if r != 2.0 {
				t.Fatalf("expected 2.0, got %v", r)
			}
		})
	})

	t.Run("Reflect", func(t *testing.T) {
		e := expr.Build[bool](func(arg *TestType) expr.Bool {
			return expr.Ref(&arg.Bool)
		})
		et, r := e.Reflect()
		if et != expr.TypeRef {
			t.Fatalf("expected expr type %v, got %v", expr.TypeRef, et)
		}
		er := [3]any{0, reflect.TypeOf((*TestType)(nil)).Elem(), []string{"Bool"}}
		if !reflect.DeepEqual(r, er) {
			t.Fatalf("expected expr values %v, got %v", er, r)
		}
	})
}
