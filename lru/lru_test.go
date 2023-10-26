package lru

import "testing"

func TestLRU(t *testing.T) {
	t.Run("Testing Put()", func(t *testing.T) {
		lru := NewLRU[int](3)
		lru.Put("one", 1)
		lru.Put("two", 2)
		lru.Put("three", 3)

		got, ok := lru.storage["one"]
		if !ok {
			t.Error("Not Okay")
		}
		if got.GetVal().value != 1 {
			t.Errorf("got: %v, expected: %v", got.GetVal().value, 1)
		}

		got2 := lru.dl.GetHead()
		if got2.GetVal().value != 3 {
			t.Errorf("got: %v, expected: %v", got2.GetVal().value, 3)
		}

		got3 := got2.GetNext()
		if got3.GetVal().value != 2 {
			t.Errorf("got: %v, expected: %v", got3.GetVal().value, 2)
		}

		lru.Put("four", 4)
		got4 := lru.dl.GetHead()
		if got4.GetVal().value != 4 {
			t.Errorf("got: %v, expected: %v", got4.GetVal().value, 4)
		}
		_, ok2 := lru.storage["one"]
		if ok2 != false {
			t.Error("Not Okay")
		}
		got5 := got4.GetNext()
		if got5.GetVal().value != 3 {
			t.Errorf("got: %v, expected: %v", got5.GetVal().value, 3)
		}
	})
}
