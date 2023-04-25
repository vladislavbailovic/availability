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
