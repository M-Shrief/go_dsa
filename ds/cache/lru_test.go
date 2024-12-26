package cache

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
		if ok2 {
			t.Error("Not Okay")
		}
		got5 := got4.GetNext()
		if got5.GetVal().value != 3 {
			t.Errorf("got: %v, expected: %v", got5.GetVal().value, 3)
		}

		lru.Put("five", 5)
		got6 := lru.dl.GetHead()
		if got6.GetVal().value != 5 {
			t.Errorf("got: %v, expected: %v", got4.GetVal().value, 4)
		}
		_, ok3 := lru.storage["two"]
		if ok3 {
			t.Error("Not Okay")
		}
		got7 := got6.GetNext()
		if got7.GetVal().value != 4 {
			t.Errorf("got: %v, expected: %v", got5.GetVal().value, 4)
		}

		_, exist := lru.Get("two")
		if exist {
			t.Error("Not Okay")
		}
		got8, exist := lru.Get("three")
		if !exist {
			t.Error("Not Okay")
		}
		if got8 != 3 {
			t.Errorf("got: %v, expected: %v", got8, 3)
		}

		got9, exist := lru.Get("four")
		if !exist {
			t.Error("Not Okay")
		}
		if got9 != 4 {
			t.Errorf("got: %v, expected: %v", got9, 4)
		}
	})

	t.Run("Testing Get()", func(t *testing.T) {
		lru := NewLRU[int](3)
		lru.Put("one", 1)
		lru.Put("two", 2)
		lru.Put("three", 3)

		got, ok := lru.Get("one")
		if ok == false {
			t.Error("Not Okay")
		}
		if got != 1 {
			t.Errorf("got: %v, expected: %v", got, 1)
		}
		got2 := lru.dl.GetHead()
		if got2.GetVal().value != 1 {
			t.Errorf("got: %v, expected: %v", got2.GetVal().value, 1)
		}
		got3 := lru.dl.GetTail()
		if got3.GetVal().value != 2 {
			t.Errorf("got: %v, expected: %v", got3.GetVal().value, 2)
		}
		_, ok2 := lru.Get("four")
		if ok2 {
			t.Error("Not Okay")
		}

		lru.Put("four", 4)

		got4 := lru.dl.GetHead()
		if got4.GetVal().value != 4 {
			t.Errorf("got: %v, expected: %v", got2.GetVal().value, 1)
		}
		got5 := lru.dl.GetTail()
		if got5.GetVal().value != 3 {
			t.Errorf("got: %v, expected: %v", got3.GetVal().value, 2)
		}

		got6, ok3 := lru.Get("four")
		if !ok3 {
			t.Error("Not Okay")
		}
		if got6 != 4 {
			t.Errorf("got: %v, expected: %v", got6, 4)
		}

		_, ok4 := lru.Get("two")
		if ok4 {
			t.Error("Not Okay")
		}
	})
}
