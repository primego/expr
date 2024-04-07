package expr

import "sync"

// ContextStore is a store of contexts. Contexts can be added to and deleted from it. It is safe for concurrent use.
type contextStore struct {
	data  map[*context]struct{}
	mutex sync.Mutex
}

// Contexts is a global store of contexts. It is used to globally track and find arguments or fields of arguments.
var contexts = &contextStore{
	data: make(map[*context]struct{}),
}

func (s *contextStore) add(ctx *context) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[ctx] = struct{}{}
}

func (s *contextStore) delete(ctx *context) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, ctx)
}
