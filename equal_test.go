package expr_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/primego/expr"
)

func TestEqual(t *testing.T) {
	t.Run("Eval", func(t *testing.T) {
		t.Run("bool values", func(t *testing.T) {
			e := expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(true),
					expr.Value(true),
				)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected true == true to be true, got false")
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(false),
					expr.Value(false),
				)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected false == false to be true, got false")
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(true),
					expr.Value(false),
				)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected true == false to be false, got true")
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(false),
					expr.Value(true),
				)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected false == true to be false, got true")
			}
		})

		t.Run("bool ref and value", func(t *testing.T) {
			type TestType struct {
				Value bool
			}

			e := expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(true),
				)
			})
			r, err := e.Eval(&TestType{Value: true})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected true == true to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: false})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected false == true to be false, got true")
			}

			e = expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(false),
				)
			})
			r, err = e.Eval(&TestType{Value: false})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected false == false to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: true})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected true == false to be false, got true")
			}
		})

		t.Run("bool refs", func(t *testing.T) {
			type TestType1 struct {
				Value bool
			}
			type TestType2 struct {
				Value bool
			}

			e := expr.Build[bool](func(test1 *TestType1, test2 *TestType2) expr.Bool {
				return expr.Equal(
					expr.Ref(&test1.Value),
					expr.Ref(&test2.Value),
				)
			})
			r, err := e.Eval(&TestType1{Value: true}, &TestType2{Value: true})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected true == true to be true, got false")
			}
			r, err = e.Eval(&TestType1{Value: true}, &TestType2{Value: false})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected true == false to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: false}, &TestType2{Value: true})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected false == true to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: false}, &TestType2{Value: false})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected false == false to be true, got false")
			}
		})

		t.Run("float64 values", func(t *testing.T) {
			e := expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(1.0),
					expr.Value(1.0),
				)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 1.0 == 1.0 to be true, got false")
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(0.0),
					expr.Value(1.0),
				)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 0.0 == 1.0 to be false, got true")
			}
		})

		t.Run("float64 ref and value", func(t *testing.T) {
			type TestType struct {
				Value float64
			}

			e := expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(1.0),
				)
			})
			r, err := e.Eval(&TestType{Value: 1.0})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 1.0 == 1.0 to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: 0.0})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 0.0 == 1.0 to be false, got true")
			}

			e = expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(0.0),
				)
			})
			r, err = e.Eval(&TestType{Value: 0.0})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 0.0 == 0.0 to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: 1.0})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 1.0 == 0.0 to be false, got true")
			}
		})

		t.Run("float64 refs", func(t *testing.T) {
			type TestType1 struct {
				Value float64
			}
			type TestType2 struct {
				Value float64
			}

			e := expr.Build[bool](func(test1 *TestType1, test2 *TestType2) expr.Bool {
				return expr.Equal(
					expr.Ref(&test1.Value),
					expr.Ref(&test2.Value),
				)
			})
			r, err := e.Eval(&TestType1{Value: 1.0}, &TestType2{Value: 1.0})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 1.0 == 1.0 to be true, got false")
			}
			r, err = e.Eval(&TestType1{Value: 1.0}, &TestType2{Value: 0.0})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 1.0 == 0.0 to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: 0.0}, &TestType2{Value: 1.0})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 0.0 == 1.0 to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: 0.0}, &TestType2{Value: 0.0})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 0.0 == 0.0 to be true, got false")
			}
		})

		t.Run("int values", func(t *testing.T) {
			e := expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(1),
					expr.Value(1),
				)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 1 == 1 to be true, got false")
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(0),
					expr.Value(1),
				)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 0 == 1 to be false, got true")
			}
		})

		t.Run("int ref and value", func(t *testing.T) {
			type TestType struct {
				Value int
			}

			e := expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(1),
				)
			})
			r, err := e.Eval(&TestType{Value: 1})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 1 == 1 to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: 0})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 0 == 1 to be false, got true")
			}

			e = expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(0),
				)
			})
			r, err = e.Eval(&TestType{Value: 0})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 0 == 0 to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: 1})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 1 == 0 to be false, got true")
			}
		})

		t.Run("int refs", func(t *testing.T) {
			type TestType1 struct {
				Value int
			}
			type TestType2 struct {
				Value int
			}

			e := expr.Build[bool](func(test1 *TestType1, test2 *TestType2) expr.Bool {
				return expr.Equal(
					expr.Ref(&test1.Value),
					expr.Ref(&test2.Value),
				)
			})
			r, err := e.Eval(&TestType1{Value: 1}, &TestType2{Value: 1})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 1 == 1 to be true, got false")
			}
			r, err = e.Eval(&TestType1{Value: 1}, &TestType2{Value: 0})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 1 == 0 to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: 0}, &TestType2{Value: 1})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected 0 == 1 to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: 0}, &TestType2{Value: 0})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected 0 == 0 to be true, got false")
			}
		})

		t.Run("string values", func(t *testing.T) {
			e := expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value("foo"),
					expr.Value("foo"),
				)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal(`expected "foo" == "foo" to be true, got false`)
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value("foo"),
					expr.Value("bar"),
				)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal(`expected "foo" == "bar" to be false, got true`)
			}
		})

		t.Run("string ref and value", func(t *testing.T) {
			type TestType struct {
				Value string
			}

			e := expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value("foo"),
				)
			})
			r, err := e.Eval(&TestType{Value: "foo"})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal(`expected "foo" == "foo" to be true, got false`)
			}
			r, err = e.Eval(&TestType{Value: "bar"})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal(`expected "bar" == "foo" to be false, got true`)
			}

			e = expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value("bar"),
				)
			})
			r, err = e.Eval(&TestType{Value: "bar"})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal(`expected "bar" == "bar" to be true, got false`)
			}
			r, err = e.Eval(&TestType{Value: "foo"})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal(`expected "foo" == "bar" to be false, got true`)
			}
		})

		t.Run("string refs", func(t *testing.T) {
			type TestType1 struct {
				Value string
			}
			type TestType2 struct {
				Value string
			}

			e := expr.Build[bool](func(test1 *TestType1, test2 *TestType2) expr.Bool {
				return expr.Equal(
					expr.Ref(&test1.Value),
					expr.Ref(&test2.Value),
				)
			})
			r, err := e.Eval(&TestType1{Value: "foo"}, &TestType2{Value: "foo"})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal(`expected "foo" == "foo" to be true, got false`)
			}
			r, err = e.Eval(&TestType1{Value: "foo"}, &TestType2{Value: "bar"})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal(`expected "foo" == "bar" to be false, got true`)
			}
			r, err = e.Eval(&TestType1{Value: "bar"}, &TestType2{Value: "foo"})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal(`expected "bar" == "foo" to be false, got true`)
			}
			r, err = e.Eval(&TestType1{Value: "bar"}, &TestType2{Value: "bar"})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal(`expected "bar" == "bar" to be true, got false`)
			}
		})

		t.Run("time values", func(t *testing.T) {
			testTime := time.Now()

			e := expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(testTime),
					expr.Value(testTime),
				)
			})
			r, err := e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected testTime == testTime to be true, got false")
			}

			e = expr.Build[bool](func() expr.Bool {
				return expr.Equal(
					expr.Value(testTime),
					expr.Value(testTime.Add(1)),
				)
			})
			r, err = e.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected testTime == testTime.Add(1) to be false, got true")
			}
		})

		t.Run("time ref and value", func(t *testing.T) {
			type TestType struct {
				Value time.Time
			}

			testTime := time.Now()

			e := expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(testTime),
				)
			})
			r, err := e.Eval(&TestType{Value: testTime})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected testTime == testTime to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: testTime.Add(1)})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected testTime == testTime.Add(1) to be false, got true")
			}

			e = expr.Build[bool](func(test *TestType) expr.Bool {
				return expr.Equal(
					expr.Ref(&test.Value),
					expr.Value(testTime.Add(1)),
				)
			})
			r, err = e.Eval(&TestType{Value: testTime.Add(1)})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected testTime.Add(1) == testTime.Add(1) to be true, got false")
			}
			r, err = e.Eval(&TestType{Value: testTime})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected testTime == testTime.Add(1) to be false, got true")
			}
		})

		t.Run("time refs", func(t *testing.T) {
			type TestType1 struct {
				Value time.Time
			}
			type TestType2 struct {
				Value time.Time
			}

			testTime := time.Now()

			e := expr.Build[bool](func(test1 *TestType1, test2 *TestType2) expr.Bool {
				return expr.Equal(
					expr.Ref(&test1.Value),
					expr.Ref(&test2.Value),
				)
			})
			r, err := e.Eval(&TestType1{Value: testTime}, &TestType2{Value: testTime})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected testTime == testTime to be true, got false")
			}
			r, err = e.Eval(&TestType1{Value: testTime}, &TestType2{Value: testTime.Add(1)})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected testTime == testTime.Add(1) to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: testTime.Add(1)}, &TestType2{Value: testTime})
			if err != nil {
				t.Fatal(err)
			}
			if r {
				t.Fatal("expected testTime.Add(1) == testTime to be false, got true")
			}
			r, err = e.Eval(&TestType1{Value: testTime.Add(1)}, &TestType2{Value: testTime.Add(1)})
			if err != nil {
				t.Fatal(err)
			}
			if !r {
				t.Fatal("expected testTime.Add(1) == testTime.Add(1) to be true, got false")
			}
		})
	})

	t.Run("Reflect", func(t *testing.T) {
		type TestType struct {
			Value int
		}

		e := expr.Build[bool](func(arg *TestType) expr.Bool {
			return expr.Equal(
				expr.Ref(&arg.Value),
				expr.Value(1),
			)
		})

		et, r := e.Reflect()
		if et != expr.TypeEqual {
			t.Fatalf("expected expr type %v, got %v", expr.TypeEqual, et)
		}
		rr, ok := r.([2]expr.Int)
		if !ok {
			t.Fatalf("expected reflect values of type [2]expr.Expr[int], got %v", reflect.TypeOf(r))
		}
		et0, _ := rr[0].Reflect()
		if et0 != expr.TypeRef {
			t.Fatalf("expected expr type %v, got %v", expr.TypeRef, et0)
		}
		et1, _ := rr[1].Reflect()
		if et1 != expr.TypeValue {
			t.Fatalf("expected expr type %v, got %v", expr.TypeValue, et1)
		}
	})
}
