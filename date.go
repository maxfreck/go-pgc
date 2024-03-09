package pgc

import (
	"fmt"
	"time"
)

type Date uint32
type Year int
type Month int
type Day int
type Era int

const MAX_YEAR = 0x3FFFFF // 4_194_303 – maximum year value is 4_194_303 BCE or 4_194_303 CE
const (
	EraBCE Era = -1
	EraCE  Era = 1
)

const year_mask = 0x7FFFFF  //8_388_607
const ce_start = 0x80000021 // Common era start 01.01.01

func MkDate(day Day, month Month, year Year) (Date, error) {
	var err = AssertDate(day, month, year)
	if err != nil {
		return 0, err
	}

	return Date(uint(day&0x1F) | uint((month&0x0F)<<5) | uint(((year+MAX_YEAR)&year_mask)<<9)), nil
}

func MkDateISO(day Day, month Month, year Year) (Date, error) {
	if year < 0 {
		year -= 1
	}

	return MkDate(day, month, year)
}

func CurrentDateUTC() (Date, error) {
	now := time.Now().UTC()
	return MkDate(Day(now.Day()), Month(now.Month()), Year(now.Year()))
}

func CurrentDateLocal() (Date, error) {
	now := time.Now()
	return MkDate(Day(now.Day()), Month(now.Month()), Year(now.Year()))
}

func AssertDate(day Day, month Month, year Year) error {
	if year < -MAX_YEAR || year == 0 || year > MAX_YEAR {
		return fmt.Errorf("The absolute year value %v is out of bounds [1..%v]", year, MAX_YEAR)
	}
	if month < 1 || month > 12 {
		return fmt.Errorf("The month value %v is out of bounds [1..12]", month)
	}

	var maxDay = DaysInMonth(month, year)
	if day < 1 || day > maxDay {
		return fmt.Errorf("The day value %v is out of bounds [1..%v]", day, maxDay)
	}

	return nil
}

func DaysInMonth(month Month, year Year) Day {
	switch month {
	case 4:
		fallthrough
	case 6:
		fallthrough
	case 9:
		fallthrough
	case 11:
		return 30
	case 2:
		if year.IsLeap() {
			return 29
		} else {
			return 28
		}
	default:
		return 31
	}
}

func (year Year) IsLeap() bool {
	if year <= 0 {
		year++
	}

	if year%4 != 0 {
		return false
	}
	if year%100 != 0 {
		return true
	}
	if year%400 != 0 {
		return false
	}
	return true
}

func (d Date) Day() Day {
	return Day(d & 0x1F)
}

func (d Date) Month() Month {
	return Month((d >> 5) & 0x0F)
}

func (d Date) Year() Year {
	return Year((int(d>>9) & year_mask) - MAX_YEAR)
}

func (d Date) IsoYear() Year {
	var year = d.Year()
	if year < 0 {
		return year + 1
	}
	return year
}

func (d Date) HolocenYear() Year {
	return d.IsoYear() + 10000
}

func DaysBetween(d1 Date, d2 Date) uint {
	var dateMin = min(d1, d2)
	var dateMax = max(d1, d2)

	return g(dateMax.Day(), dateMax.Month(), dateMax.IsoYear()) - g(dateMin.Day(), dateMin.Month(), dateMin.IsoYear())
}

func g(d Day, m Month, y Year) uint {
	m = (m + 9) % 12
	y = y - Year(m/10)
	return uint(365*y) + uint(y/4) - uint(y/100) + uint(y/400) + uint(m*306+5)/10 + uint(d-1)
}

func (d Date) Era() Era {
	if d < ce_start {
		return EraBCE
	}
	return EraCE
}

func (d Date) NextDay() (Date, error) {
	var day = d.Day() + 1
	var month = d.Month()
	var year = d.Year()

	if day > DaysInMonth(month, year) {
		day = 1
		month += 1
		if month > 12 {
			month = 1
			year += 1
			if year == 0 {
				year += 1
			}
		}
	}

	return MkDate(day, month, year)
}

func (d Date) PrevDay() (Date, error) {
	var day = d.Day() - 1
	var month = d.Month()
	var year = d.Year()

	if day < 1 {
		month -= 1
		if month < 1 {
			month = 12
			year -= 1
			if year == 0 {
				year -= 1
			}
		}
		day = DaysInMonth(month, year)
	}

	return MkDate(day, month, year)
}
