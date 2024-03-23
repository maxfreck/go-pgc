package pgc

import "testing"

func TestBasicTime(t *testing.T) {
	var time, _ = MkTime(2, 19, 15, 480)

	var hour = time.Hour()
	if hour != 2 {
		t.Errorf("Hour expected 02 got %v", hour)
	}

	var minute = time.Minute()
	if minute != 19 {
		t.Errorf("Minute expected 19 got %v", minute)
	}

	var second = time.Second()
	if second != 15 {
		t.Errorf("Second expected 15 got %v", second)
	}

	var split = time.Split()
	if split != 480 {
		t.Errorf("Split expected 480 got %v", split)
	}

}
