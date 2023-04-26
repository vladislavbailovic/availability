package data

import "testing"

func Test_IntArgAt(t *testing.T) {
	func(empty ...any) {
		if IntArgAt(empty, 161) != 0 {
			t.Error("expected zero")
		}
	}()

	func(one ...any) {
		if IntArgAt(one, 161) != 0 {
			t.Error("expected zero")
		}
	}(1)

	func(test ...any) {
		if IntArgAt(test, 0) != 161 {
			t.Error("expected first")
		}
		if IntArgAt(test, 1) != 13 {
			t.Error("expected second")
		}
		if IntArgAt(test, 2) != 12 {
			t.Error("expected third")
		}
	}(161, 13, 12)
}

func Test_DataID_ToIDs(t *testing.T) {
	var want1 int32 = 1312
	x := DataID(1312)
	if x.ToItemID() != want1 {
		t.Error("item id conversion failed")
	}

	var want2 int = 1312
	if x.ToNumericID() != want2 {
		t.Error("numeric id conversion failed")
	}
}
