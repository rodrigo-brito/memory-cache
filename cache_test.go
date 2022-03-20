package cache

import (
	"reflect"
	"testing"
	"time"
)

func equal[T any](t *testing.T, value, expect T) {
	t.Helper()
	if !reflect.DeepEqual(value, expect) {
		t.Errorf("Expected %v, got %v", expect, value)
	}
}

func TestCache(t *testing.T) {
	cache := New[int](time.Millisecond)
	cache.Set("one", 1, time.Minute)
	cache.Set("two", 2, time.Minute)
	cache.Set("expired", 1, 0)

	t.Run("found", func(t *testing.T) {
		value, ok := cache.Get("one")
		equal(t, ok, true)
		equal(t, value, 1)
	})

	t.Run("not found", func(t *testing.T) {
		value, ok := cache.Get("not found")
		equal(t, ok, false)
		equal(t, value, 0)
	})

	t.Run("expired TTL", func(t *testing.T) {
		value, ok := cache.Get("expired")
		equal(t, ok, false)
		equal(t, value, 0)
	})

	t.Run("get multi", func(t *testing.T) {
		values, ok := cache.GetMulti("one", "two")
		equal(t, ok, true)
		equal(t, values, []int{1, 2})
	})
}
