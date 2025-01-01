package cache

import "testing"

func TestMFU(t *testing.T) {
	t.Run("Testing Put()", func(t *testing.T) {
		mfu := NewMFU[int](3)
		mfu.Put("one", 1)
		mfu.Put("two", 2)
		mfu.Put("three", 3)

		if mfu.size != 3 {
			t.Errorf("got: %v, expected: %v", mfu.size, 3)
		}

		got, ok := mfu.storage["one"]
		if !ok {
			t.Error("Not Okay")
		}
		if got.val != 1 {
			t.Errorf("got: %v, expected: %v", got.val, 1)
		}

		mfu.Get("one")
		mfu.Get("one")
		mfu.Get("two")
		mfu.Get("two")
		mfu.Get("two")
		mfu.Get("three")
		// // Score: one.freq= 2, two.freq= 3, three.freq=1

		top := mfu.GetTop()
		if top.val != 2 {
			t.Errorf("got: %v, expected: %v", top.val, 2)
		}

		mfu.Put("four", 4)

		got2, ok := mfu.storage["four"]
		if !ok {
			t.Error("Not Okay")
		}
		if got2.val != 4 {
			t.Errorf("got: %v, expected: %v", got2.val, 4)
		}

		_, ok = mfu.storage["two"]
		if ok {
			t.Error("Shouldn't exist")
		}

		mfu.Get("one")
		mfu.Get("one")
		mfu.Get("three")
		mfu.Get("four")
		mfu.Get("four")
		mfu.Get("four")
		// Freqs -> one.freq = 4, three.freq = 2, four.freq = 3;

		top = mfu.GetTop()
		if top.val != 1 {
			t.Errorf("got: %v, expected: %v", top.val, 1)
		}

		mfu.Put("five", 5)
		got3, ok := mfu.storage["five"]
		if !ok {
			t.Error("Not Okay")
		}
		if got3.val != 5 {
			t.Errorf("got: %v, expected: %v", got3.val, 5)
		}
		_, ok = mfu.storage["one"]
		if ok {
			t.Error("Shouldn't exist")
		}
	})

	t.Run("Testing Get()", func(t *testing.T) {
		mfu := NewMFU[int](3)
		mfu.Put("one", 1)
		mfu.Put("two", 2)
		mfu.Put("three", 3)

		got, ok := mfu.Get("one")
		if ok == false {
			t.Error("Not Okay")
		}
		if got != 1 {
			t.Errorf("got: %v, expected: %v", got, 1)
		}

		got2, ok := mfu.Get("two")
		if ok == false {
			t.Error("Not Okay")
		}
		if got2 != 2 {
			t.Errorf("got: %v, expected: %v", got2, 2)
		}

		mfu.Get("one")
		mfu.Get("two")
		mfu.Get("two")
		mfu.Get("three")
		// // Score: one.freq= 2, two.freq= 3, three.freq=1

		got3 := mfu.maxHeap.Top()
		if got3.val != 2 {
			t.Errorf("got: %v, expected: %v", got3.val, 2)
		}

		// getting non-existing value
		_, ok2 := mfu.Get("four")
		if ok2 {
			t.Error("Not Okay")
		}

		mfu.Put("four", 4)

		got4, ok := mfu.Get("four")
		if ok == false {
			t.Error("Not Okay")
		}
		if got4 != 4 {
			t.Errorf("got: %v, expected: %v", got2, 4)
		}

		got5 := mfu.maxHeap.Top()
		if got5.val != 1 {
			t.Errorf("got: %v, expected: %v", got5.val, 1)
		}
		// Evict least frequently used correctly
		_, ok2 = mfu.Get("two")
		if ok2 {
			t.Error("Not Okay")
		}

		mfu.Get("one")
		mfu.Get("one")
		mfu.Get("three")
		mfu.Get("four")
		mfu.Get("four")
		// Freqs -> one.freq = 4, three.freq = 2, four.freq = 3;

		mfu.Put("five", 5)

		got6, ok := mfu.Get("five")
		if ok == false {
			t.Error("Not Okay")
		}
		if got6 != 5 {
			t.Errorf("got: %v, expected: %v", got6, 5)
		}

		got7 := mfu.maxHeap.Top()
		if got7.val != 4 {
			t.Errorf("got: %v, expected: %v", got7.val, 4)
		}
		_, ok2 = mfu.Get("one")
		if ok2 {
			t.Error("Not Okay")
		}
	})
}
