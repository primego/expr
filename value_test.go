package expr_test

import (
	"testing"
	"time"

	"github.com/primego/expr"
)

func TestValue(t *testing.T) {
	t.Run("Eval", func(t *testing.T) {
		t.Run("bool", func(t *testing.T) {
			e := expr.Build[bool](func() expr.Bool {
				return expr.Value(true)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected true, got false")
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Value(false)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected false, got true")
			}
		})

		t.Run("float64", func(t *testing.T) {
			e := expr.Build[float64](func() expr.Float64 {
				return expr.Value(1.0)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r != 1.0 {
				t.Fatalf("expected 1.0, got %v", r)
			}

			e = expr.Build[float64](func() expr.Float64 {
				return expr.Value(2.0)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r != 2.0 {
				t.Fatalf("expected 2.0, got %v", r)
			}
		})

		t.Run("int", func(t *testing.T) {
			e := expr.Build[int](func() expr.Int {
				return expr.Value(1)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r != 1 {
				t.Fatalf("expected 1, got %v", r)
			}

			e = expr.Build[int](func() expr.Int {
				return expr.Value(2)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r != 2 {
				t.Fatalf("expected 2, got %v", r)
			}
		})

		t.Run("string", func(t *testing.T) {
			e := expr.Build[string](func() expr.String {
				return expr.Value("foo")
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r != "foo" {
				t.Fatalf(`expected "foo", got "%v"`, r)
			}

			e = expr.Build[string](func() expr.String {
				return expr.Value("bar")
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r != "bar" {
				t.Fatalf(`expected "bar", got "%v"`, r)
			}
		})

		t.Run("time", func(t *testing.T) {
			e := expr.Build[time.Time](func() expr.Time {
				return expr.Value(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r.Equal(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)) {
				t.Fatalf("expected 2020-01-01 00:00:00 +0000 UTC, got %v", r)
			}
		})
	})

	t.Run("Reflect", func(t *testing.T) {
		e := expr.Build[int](func() expr.Int {
			return expr.Value(1)
		})
		et, r := e.Reflect()
		if et != expr.TypeValue {
			t.Fatalf("expected expr type %v, got %v", expr.TypeValue, et)
		}
		if r != 1 {
			t.Fatalf("expected expr value 1, got %v", r)
		}
	})
}
