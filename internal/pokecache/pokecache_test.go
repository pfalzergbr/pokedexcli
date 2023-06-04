package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)

	cases := []struct {
		inputKey   string
		inputValue []byte
	}{
		{inputKey: "pikachu",
			inputValue: []byte("pikachu")},
	}

	for _, cas := range cases {

		cache.Add(cas.inputKey, cas.inputValue)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("expected to find %s in cache", string(cas.inputValue))
			continue
		}

		if string(actual) != string(cas.inputValue) {
			t.Errorf("expected %s, got %v", string(cas.inputValue), actual)
			continue
		}
	}
}
func TestReap (t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))

	time.Sleep(interval + 1)

	_, ok := cache.Get(keyOne)

	if ok {
		t.Errorf("expected %s to be reaped", keyOne)
	}


}
