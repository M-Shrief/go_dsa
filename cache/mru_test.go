package cache

import "testing"

func TestMRU(t *testing.T) {
	t.Run("Testing Put()", func(t *testing.T) {
		mru := NewMRU[int](3)
		mru.Put("one", 1)
		mru.Put("two", 2)
		mru.Put("three", 3)

		got, ok := mru.storage["one"]
		if !ok {
			t.Error("Not Okay")
		}
		if got.GetVal().value != 1 {
			t.Errorf("got: %v, expected: %v", got.GetVal().value, 1)
		}

		got2 := mru.dl.GetHead()
		if got2.GetVal().value != 3 {
			t.Errorf("got: %v, expected: %v", got2.GetVal().value, 3)
		}

		got3 := got2.GetNext()
		if got3.GetVal().value != 2 {
			t.Errorf("got: %v, expected: %v", got3.GetVal().value, 2)
		}

		mru.Put("four", 4)
		got4 := mru.dl.GetHead()
		if got4.GetVal().value != 4 {
			t.Errorf("got: %v, expected: %v", got4.GetVal().value, 4)
		}
		_, ok2 := mru.storage["one"]
		if ok2 {
			t.Error("Not Okay")
		}
		got5 := got4.GetNext()
		if got5.GetVal().value != 3 {
			t.Errorf("got: %v, expected: %v", got5.GetVal().value, 3)
		}

		mru.Put("five", 5)
		got6 := mru.dl.GetHead()
		if got6.GetVal().value != 5 {
			t.Errorf("got: %v, expected: %v", got6.GetVal().value, 5)
		}
		_, ok3 := mru.storage["two"]
		if ok3 {
			t.Error("Not Okay")
		}
		got7 := got6.GetNext()
		if got7.GetVal().value != 4 {
			t.Errorf("got: %v, expected: %v", got7.GetVal().value, 4)
		}
	})

	t.Run("Testing Get()", func(t *testing.T) {
		mru := NewMRU[int](3)
		mru.Put("one", 1)
		mru.Put("two", 2)
		mru.Put("three", 3)

		got, ok := mru.Get("two")
		if ok == false {
			t.Error("Not Okay")
		}
		if got != 2 {
			t.Errorf("got: %v, expected: %v", got, 2)
		}

		got2 := mru.dl.GetTail()
		if got2.GetVal().value != 2 {
			t.Errorf("got: %v, expected: %v", got2.GetVal().value, 2)
		}

		_, ok2 := mru.Get("four")
		if ok2 {
			t.Error("Not Okay")
		}

		mru.Put("four", 4)

		got3, ok3 := mru.Get("four")
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != 4 {
			t.Errorf("got: %v, expected: %v", got3, 4)
		}

		got4 := mru.dl.GetTail()
		if got4.GetVal().value != 4 {
			t.Errorf("got: %v, expected: %v", got4.GetVal().value, 4)
		}
		_, ok4 := mru.Get("two")
		if ok4 {
			t.Error("Not Okay")
		}
	})
}
