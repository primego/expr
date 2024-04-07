package expr

import (
	"reflect"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
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

	type Bundle struct {
		Book1     Book
		Book2     Book
		CreatedBy User
	}

	ctx1 := &context{}
	contexts.add(ctx1)
	defer contexts.delete(ctx1)
	value1 := true
	ctx1.add(&value1)
	user1 := User{}
	ctx1.add(&user1)
	book1 := Book{}
	ctx1.add(&book1)
	bundle1 := Bundle{}
	ctx1.add(&bundle1)

	ctx2 := &context{}
	contexts.add(ctx2)
	defer contexts.delete(ctx2)
	book2 := Book{}
	ctx2.add(&book2)
	bundle2 := Bundle{}
	ctx2.add(&bundle2)
	value2 := 2
	ctx2.add(&value2)
	user2 := User{}
	ctx2.add(&user2)

	t.Run("ctx1.find with arguments or fields of arguments in ctx1", func(t *testing.T) {
		argIndex, argType, fieldPath := ctx1.find(&value1)
		if argIndex != 0 || argType != reflect.TypeOf(value1) || len(fieldPath) != 0 {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx1.find(&user1.Id)
		if argIndex != 1 || argType != reflect.TypeOf(user1) || !reflect.DeepEqual(fieldPath, []string{"Id"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx1.find(&book1.Weight)
		if argIndex != 2 || argType != reflect.TypeOf(book1) || !reflect.DeepEqual(fieldPath, []string{"Weight"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx1.find(&bundle1.Book2.Title)
		if argIndex != 3 || argType != reflect.TypeOf(bundle1) ||
			!reflect.DeepEqual(fieldPath, []string{"Book2", "Title"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}
	})

	t.Run("ctx2.find with arguments or fields of arguments in ctx2", func(t *testing.T) {
		argIndex, argType, fieldPath := ctx2.find(&book2.BorrowedAt)
		if argIndex != 0 || argType != reflect.TypeOf(book2) || !reflect.DeepEqual(fieldPath, []string{"BorrowedAt"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx2.find(&bundle2.Book1.Lost)
		if argIndex != 1 || argType != reflect.TypeOf(bundle2) ||
			!reflect.DeepEqual(fieldPath, []string{"Book1", "Lost"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx2.find(&value2)
		if argIndex != 2 || argType != reflect.TypeOf(value2) || len(fieldPath) != 0 {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx2.find(&user2.Name)
		if argIndex != 3 || argType != reflect.TypeOf(user2) || !reflect.DeepEqual(fieldPath, []string{"Name"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}
	})

	t.Run("ctx1.find with arguments or fields of arguments in ctx2", func(t *testing.T) {
		argIndex, argType, fieldPath := ctx1.find(&book2.BorrowedAt)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx1.find(&bundle2.Book1.Lost)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx1.find(&value2)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx1.find(&user2.Name)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}
	})

	t.Run("ctx2.find with arguments or fields of arguments in ctx1", func(t *testing.T) {
		argIndex, argType, fieldPath := ctx2.find(&value1)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx2.find(&user1.Id)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx2.find(&book1.Weight)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = ctx2.find(&bundle1.Book2.Title)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}
	})

	t.Run("global find with arguments or fields of arguments in ctx1 and ctx2", func(t *testing.T) {
		argIndex, argType, fieldPath := find(&value1)
		if argIndex != 0 || argType != reflect.TypeOf(value1) || len(fieldPath) != 0 {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&user1.Id)
		if argIndex != 1 || argType != reflect.TypeOf(user1) || !reflect.DeepEqual(fieldPath, []string{"Id"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&book1.Weight)
		if argIndex != 2 || argType != reflect.TypeOf(book1) || !reflect.DeepEqual(fieldPath, []string{"Weight"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&bundle1.Book2.Title)
		if argIndex != 3 || argType != reflect.TypeOf(bundle1) ||
			!reflect.DeepEqual(fieldPath, []string{"Book2", "Title"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&book2.BorrowedAt)
		if argIndex != 0 || argType != reflect.TypeOf(book2) || !reflect.DeepEqual(fieldPath, []string{"BorrowedAt"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&bundle2.Book1.Lost)
		if argIndex != 1 || argType != reflect.TypeOf(bundle2) ||
			!reflect.DeepEqual(fieldPath, []string{"Book1", "Lost"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&value2)
		if argIndex != 2 || argType != reflect.TypeOf(value2) || len(fieldPath) != 0 {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&user2.Name)
		if argIndex != 3 || argType != reflect.TypeOf(user2) || !reflect.DeepEqual(fieldPath, []string{"Name"}) {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}
	})

	t.Run("global find with arguments or fields of arguments not added in any tracked context", func(t *testing.T) {
		user3 := &User{}
		book3 := &Book{}
		value3 := "foo"

		// ctx3 is not added to the context store
		ctx3 := &context{}
		ctx3.add(book3)

		argIndex, argType, fieldPath := find(&user3.Id)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&book3.Title)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}

		argIndex, argType, fieldPath = find(&value3)
		if argIndex != -1 || argType != nil || fieldPath != nil {
			t.Errorf("context.find returned unexpected result: %v, %v, %v", argIndex, argType, fieldPath)
		}
	})
}
