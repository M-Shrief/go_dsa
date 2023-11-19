package hashmap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	t.Run("Testing Push()", func(t *testing.T) {
		hm := New[int](3)

		hash, ok := hm.Put("one", 1)
		if ok != nil {
			t.Error("Not Okay")
		}

		if hm.table[hash].value != 1 {
			t.Errorf("Got: %v, Want: %v", hm.table[hash].value, 1)
		}

		hash1, ok1 := hm.Put("one", 11)
		if ok1 != nil {
			t.Error("Not Okay")
		}

		if hm.table[hash1].value != 11 {
			t.Errorf("Got: %v, Want: %v", hm.table[hash1].value, 11)
		}

		if hm.table[hash].next.value != 1 {
			t.Errorf("Got: %v, Want: %v", hm.table[hash].next.value, 1)
		}

		hash2, ok2 := hm.Put("two", 2)
		if ok2 != nil {
			t.Error("Not Okay")
		}

		if hm.table[hash2].value != 2 {
			t.Errorf("Got: %v, Want: %v", hm.table[hash2].value, 2)
		}

		hash3, ok3 := hm.Put("four", 4)
		if ok3 == nil || hash3 != 0 {
			t.Error("Not Okay")
		}
	})

	t.Run("Testing Get()", func(t *testing.T) {
		hm := New[int](5)

		n, err := hm.Get("one")
		if err == nil || n != 0 {
			t.Error("Not Okay")
		}

		hash, _ := hm.Put("one", 1)
		val, ok := hm.Get("one")
		if ok != nil {
			t.Error("Not Okay")
		}

		if val != 1 {
			t.Errorf("Got: %v, Want: %v", val, 1)
		}

		t.Run("Testing collisions handling with getHelper()", func(t *testing.T) {
			hm.table[hash] = &node[int]{
				key:   "eleven",
				value: 11,
				next:  hm.table[hash],
			}

			hm.table[hash] = &node[int]{
				key:   "one hundred",
				value: 100,
				next:  hm.table[hash],
			}

			val1, ok1 := hm.getHelper("one hundred", hash)
			if ok1 != nil {
				t.Error("Not Okay")
			}

			if val1 != 100 {
				t.Errorf("Got: %v, Want: %v", val1, 100)
			}

			val2, ok2 := hm.getHelper("eleven", hash)
			if ok2 != nil {
				t.Error("Not Okay")
			}

			if val2 != 11 {
				t.Errorf("Got: %v, Want: %v", val2, 11)
			}

			val3, ok3 := hm.getHelper("one", hash)
			if ok3 != nil {
				t.Error("Not Okay")
			}

			if val3 != 1 {
				t.Errorf("Got: %v, Want: %v", val3, 1)
			}

		})
	})
}
