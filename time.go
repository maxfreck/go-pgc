package pgc

import "fmt"

type Time uint32
type Hour int
type Minute int
type Second int
type Split int

func AssertTime(hour Hour, minute Minute, second Second, split Split) error {
	if hour < 0 || hour > 23 {
		return fmt.Errorf("invalid hour value: %v", hour)
	}
	if minute < 0 || minute > 59 {
		return fmt.Errorf("invalid minute value: %v", minute)
	}
	if second < 0 || second > 59 {
		return fmt.Errorf("invalid second value: %v", second)
	}
	if split < 0 || split > 32_768 {
		return fmt.Errorf("split %v is out of bound [0..32768]", split)
	}

	return nil
}
