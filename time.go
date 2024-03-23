package pgc

import "fmt"

type Time uint32
type Hour int
type Minute int
type Second int
type Split int

func MkTime(hour Hour, minute Minute, second Second, split Split) (Time, error) {
	var err = AssertTime(hour, minute, second, split)
	if err != nil {
		return 0, err
	}

	return Time(uint((hour&0x1F)<<27) | uint((minute&0x3F)<<21) | uint((second&0x3F)<<15) | uint(split&0x3FFF)), nil
}

func (t Time) Hour() Hour {
	return Hour((t >> 27) & 0x1F)
}

func (t Time) Minute() Minute {
	return Minute((t >> 21) & 0x3F)
}

func (t Time) Second() Second {
	return Second((t >> 15) & 0x3F)
}

func (t Time) Split() Split {
	return Split(t & 0x7FFF)
}

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
