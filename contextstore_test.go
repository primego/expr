package expr

import "testing"

func TestContextStore(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		store := &contextStore{
			data: make(map[*context]struct{}),
		}

		ctx := &context{}
		store.add(ctx)
		if _, ok := store.data[ctx]; !ok {
			t.Fatal("contextStore.add did not add the context to the store")
		}
	})

	t.Run("delete", func(t *testing.T) {
		store := &contextStore{
			data: make(map[*context]struct{}),
		}

		ctx := &context{}
		store.data[ctx] = struct{}{}

		store.delete(ctx)
		if _, ok := store.data[ctx]; ok {
			t.Fatal("contextStore.delete did not delete the context from the store")
		}
	})
}
