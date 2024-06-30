package cache

import "testing"

func TestLFU(t *testing.T) {
	t.Run("Testing Put()", func(t *testing.T) {
		lfu := NewLFU[int](3)
		lfu.Put("one", 1)
		lfu.Put("two", 2)
		lfu.Put("three", 3)

		if lfu.size != 3 {
			t.Errorf("got: %v, expected: %v", lfu.size, 3)
		}

		got, ok := lfu.storage["one"]
		if !ok {
			t.Error("Not Okay")
		}
		if got.val != 1 {
			t.Errorf("got: %v, expected: %v", got.val, 1)
		}

		// Increasing keys.freq, the least one is "two"
		lfu.Get("one")
		lfu.Get("one")
		lfu.Get("two")
		lfu.Get("three")
		lfu.Get("three")

		lfu.Put("four", 4)
		got2 := lfu.mh.Top()
		if got2.val != 4 {
			t.Errorf("got: %v, expected: %v", got2.val, 4)
		}

		_, notOk := lfu.storage["two"]
		if notOk {
			t.Error("Not Okay")
		}
	})

	t.Run("Testing Get()", func(t *testing.T) {
		lfu := NewLFU[int](3)
		lfu.Put("one", 1)
		lfu.Put("two", 2)
		lfu.Put("three", 3)

		got, ok := lfu.Get("one")
		if ok == false {
			t.Error("Not Okay")
		}
		if got != 1 {
			t.Errorf("got: %v, expected: %v", got, 1)
		}

		got2, ok := lfu.Get("two")
		if ok == false {
			t.Error("Not Okay")
		}
		if got2 != 2 {
			t.Errorf("got: %v, expected: %v", got2, 2)
		}

		got3 := lfu.mh.Top()
		if got3.val != 3 {
			t.Errorf("got: %v, expected: %v", got3.val, 3)
		}
		_, ok2 := lfu.Get("four")
		if ok2 {
			t.Error("Not Okay")
		}

		lfu.Put("four", 4)

		got4 := lfu.mh.Top()
		if got4.val != 4 {
			t.Errorf("got: %v, expected: %v", got4.val, 4)
		}
		// Evict least frequently used correctly
		_, ok2 = lfu.Get("three")
		if ok2 {
			t.Error("Not Okay")
		}

		lfu.Get("one")
		lfu.Get("one")
		lfu.Get("two")
		lfu.Get("four")
		lfu.Get("four")
		lfu.Get("four")
		// Now: one.freq = 3, two.freq = 2, four.freq = 3
		lfu.Put("five", 5)

		got5 := lfu.mh.Top()
		if got5.val != 5 {
			t.Errorf("got: %v, expected: %v", got5.val, 5)
		}
		// Evict least frequently used correctly
		_, ok2 = lfu.Get("two")
		if ok2 {
			t.Error("Not Okay")
		}
	})
}
