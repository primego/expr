package expr_test

import (
	"testing"
	"time"

	"github.com/primego/expr"
)

func TestBuild(t *testing.T) {
	type User struct {
		Id   int64
		Name string
	}

	type Book struct {
		Id         int64
		Title      string
		Weight     float64
		Lost       bool
		BorrowedBy int64
		BorrowedAt time.Time
	}

	t.Run("Build[bool] returning a simple value", func(t *testing.T) {
		var er expr.Bool
		ar := expr.Build[bool](func() expr.Bool {
			er = expr.Value(true)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[bool]")
		}
	})

	t.Run("Build[bool] with a simple value argument and returning the ref of the value", func(t *testing.T) {
		var er expr.Bool
		ar := expr.Build[bool](func(v *bool) expr.Bool {
			if v == nil {
				t.Fatal("builder argument is nil")
			}
			er = expr.Ref(v)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[bool]")
		}
	})

	t.Run("Build[bool] with a struct argument and returning a ref to a bool field", func(t *testing.T) {
		var er expr.Bool
		ar := expr.Build[bool](func(book *Book) expr.Bool {
			if book == nil {
				t.Fatal("builder argument is nil")
			}
			er = expr.Ref(&book.Lost)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[bool]")
		}
	})

	t.Run("Build[bool] with two struct arguments and returning the equality of two refs", func(t *testing.T) {
		var er expr.Bool
		ar := expr.Build[bool](func(user *User, book *Book) expr.Bool {
			if user == nil || book == nil {
				t.Fatal("builder argument is nil")
			}
			er = expr.EqualRR(&user.Id, &book.BorrowedBy)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[bool]")
		}
	})

	t.Run("Build[bool] with non-pointer simple value argument", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic")
			}
			if r != "builder arguments must be pointers" {
				t.Fatalf("unexpected panic: %v", r)
			}
		}()
		expr.Build[bool](func(v bool) expr.Bool {
			return expr.Ref(&v)
		})
	})

	t.Run("Build[bool] with non-pointer struct argument", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic")
			}
			if r != "builder arguments must be pointers" {
				t.Fatalf("unexpected panic: %v", r)
			}
		}()
		expr.Build[bool](func(book Book) expr.Bool {
			return expr.Ref(&book.Lost)
		})
	})

	t.Run("Build[bool] not returning an expr.Expr[bool]", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic")
			}
			if r != "builder must return an expr.Expr[bool]" {
				t.Fatalf("unexpected panic: %v", r)
			}
		}()
		expr.Build[bool](func() bool {
			return true
		})
	})

	t.Run("Build[float64] with a struct argument and returning a ref to a float64 field", func(t *testing.T) {
		var er expr.Float64
		ar := expr.Build[float64](func(book *Book) expr.Float64 {
			if book == nil {
				t.Fatal("builder argument is nil")
			}
			er = expr.Ref(&book.Weight)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[float64]")
		}
	})

	t.Run("Build[int64] with a struct argument and returning a ref to a int64 field", func(t *testing.T) {
		var er expr.Int64
		ar := expr.Build[int64](func(user *User) expr.Int64 {
			if user == nil {
				t.Fatal("builder argument is nil")
			}
			er = expr.Ref(&user.Id)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[int64]")
		}
	})

	t.Run("Build[string] with a struct argument and returning a ref to a string field", func(t *testing.T) {
		var er expr.String
		ar := expr.Build[string](func(user *User) expr.String {
			if user == nil {
				t.Fatal("builder argument is nil")
			}
			er = expr.Ref(&user.Name)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[string]")
		}
	})

	t.Run("Build[time.Time] with a struct argument and returning a ref to a time field", func(t *testing.T) {
		var er expr.Time
		ar := expr.Build[time.Time](func(book *Book) expr.Time {
			if book == nil {
				t.Fatal("builder argument is nil")
			}
			er = expr.Ref(&book.BorrowedAt)
			return er
		})
		if ar != er {
			t.Fatal("unexpected return of Build[time.Time]")
		}
	})
}
