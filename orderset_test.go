package orderset

import "testing"

var set = NewIntegerOrderSet()


func TestIntegerOrderSet_Add(t *testing.T) {
	set.Add(99)
	set.Add(88)
	set.Add(77)
}

func TestIntegerOrderSet_Del(t *testing.T) {
	set.Del(88)
}

func TestIntegerOrderSet_Contains(t *testing.T) {
	b1 := set.Contains(99)
	if !b1 {
		t.Error("Contain error")
	}

	b2 := set.Contains(88)
	if b2 {
		t.Error("Contain error")
	}
}

func TestIntegerOrderSet_ToSlice(t *testing.T) {
	s := set.ToSlice()

	count := len(s)
	if count != 2 {
		t.Error("Size error")
	}

	if s[0] != 99 {
		t.Error("Position error")
	}
	if s[1] != 77 {
		t.Error("Position error")
	}
}

