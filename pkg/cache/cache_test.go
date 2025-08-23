package cache

import (
	"testing"
	"time"
)

type TestCache struct {
	key      string
	value    interface{}
	priority int64
}

func (t *TestCache) Priority() int64 {
	return t.priority
}

func (t *TestCache) Key() string {
	return t.key
}

func (t *TestCache) Value() interface{} {
	return t.value
}

func (t *TestCache) UpdatePriority() {
	t.priority = time.Now().UnixNano()
}

func TestPutAndGet(t *testing.T) {
	cache := NewCache(5)
	testItem := &TestCache{
		key:      "test-key",
		value:    "test-value",
		priority: time.Now().UnixNano(),
	}
	cache.Put("test-key", testItem)
	result := cache.Get("test-key")
	if result == nil {
		t.Error("Expected to get item, got nil")
	}
	if result != "test-value" {
		t.Errorf("Expected test-value, got %v", result)
	}
}

func TestNoExist(t *testing.T) {
	cache := NewCache(5)

	result := cache.Get("non-existent")
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestCapacityOverflow(t *testing.T) {
	cache := NewCache(2)

	item1 := &TestCache{key: "key1", value: "value1", priority: 1}
	item2 := &TestCache{key: "key2", value: "value2", priority: 2}
	item3 := &TestCache{key: "key3", value: "value3", priority: 3}

	cache.Put("key1", item1)
	cache.Put("key2", item2)
	cache.Put("key3", item3)

	if cache.Get("key1") != nil {
		t.Error("Expected key1 to be popped")
	}
	if cache.Get("key2") == nil {
		t.Error("Expected key2 to exist")
	}
	if cache.Get("key3") == nil {
		t.Error("Expected key3 to exist")
	}
}

func TestUpdatePriority(t *testing.T) {
	cache := NewCache(5)

	item := &TestCache{
		key:      "test-key",
		value:    "test-value",
		priority: time.Now().UnixNano(),
	}

	cache.Put("test-key", item)

	originalPriority := item.priority
	time.Sleep(1 * time.Millisecond)

	cache.Get("test-key")

	if item.priority <= originalPriority {
		t.Error("Expected priority to be updated after Get operation")
	}
}
