package cache

// cache item
type item[T any] struct {
	key   string
	value T
}
