package pgc

import "testing"

func TestLeap(t *testing.T) {
	if Year(800).IsLeap() == false {
		t.Error("Is year 800 a leap year? Expected true, got false")
	}
	if Year(-1).IsLeap() == false {
		t.Error("Is year -1 a leap year? Expected true, got false")
	}
	if Year(-97).IsLeap() == false {
		t.Error("Is year -97 a leap year? Expected true, got false")
	}

	if Year(-1500).IsLeap() == true {
		t.Error("Is year -1500 a leap year? Expected false, got true")
	}
	if Year(-101).IsLeap() == true {
		t.Error("Is year -101 a leap year? Expected false, got true")
	}
}

func TestIntervals(t *testing.T) {
	const referenceDays = 1286

	var sputnikOne, _ = MkDate(4, 10, 1957)
	var gagarin, _ = MkDate(12, 4, 1961)

	var numberOfDays = DaysBetween(sputnikOne, gagarin)
	if numberOfDays != referenceDays {
		t.Errorf("Number of days between the flight of the first satellite and the first man into space. Expected %v got %v ", referenceDays, numberOfDays)
	}
}

func TestBasic(t *testing.T) {
	var platonBirth, _ = MkDate(10, 11, -427)

	var era = platonBirth.Era()
	if era != EraBCE {
		t.Errorf("Platon birth Era expected %v got %v", EraBCE, era)
	}

	var year = platonBirth.Year()
	if year != -427 {
		t.Errorf("Platon birth Year expected %v got %v", -427, year)
	}

	var isoYear = platonBirth.IsoYear()
	if isoYear != -426 {
		t.Errorf("Platon birth ISO Year expected %v got %v", -426, isoYear)
	}

	var holocenYear = platonBirth.HolocenYear()
	if holocenYear != 9574 {
		t.Errorf("Platon birth Holocen Year expected %v got %v", 9574, holocenYear)
	}

	var gagarin, _ = MkDate(12, 4, 1961)

	era = gagarin.Era()
	if era != EraCE {
		t.Errorf("Gagarin flight Era expected %v got %v", EraCE, era)
	}

	year = gagarin.Year()
	if year != 1961 {
		t.Errorf("Gagarin flight Year expected %v got %v", 1961, year)
	}

	isoYear = gagarin.IsoYear()
	if isoYear != 1961 {
		t.Errorf("Gagarin flight ISO Year expected %v got %v", 1961, year)
	}

	holocenYear = gagarin.HolocenYear()
	if holocenYear != 11961 {
		t.Errorf("Gagarin flight Year expected %v got %v", 11961, holocenYear)
	}

	var holocenEraStart, _ = MkDate(1, 1, -10000)
	holocenYear = holocenEraStart.HolocenYear()
	if holocenYear != 1 {
		t.Errorf("Holocen Era Year expected %v got %v", 1, holocenYear)
	}
}

func TestCreate(t *testing.T) {
	var gagarin, _ = MkDate(12, 4, 1961)
	var gagarinISO, _ = MkDateISO(12, 4, 1961)

	if gagarin != gagarinISO {
		t.Errorf("Gagarin flight dates don't match")
	}

	var platonBirth, _ = MkDate(10, 11, -427)
	var platonBirthISO, _ = MkDateISO(10, 11, -426)
	if platonBirth != platonBirthISO {
		t.Errorf("Platon birth dates don't match")
	}
}

func TestIteration(t *testing.T) {
	var date, _ = MkDate(31, 12, -1)
	var next, _ = date.NextDay()

	if next.Day() != 1 ||
		next.Month() != 1 ||
		next.Year() != 1 {
		t.Errorf("Iteration: next date is invalid")
	}

	var prev, _ = next.PrevDay()

	if prev != date {
		t.Errorf("Iteration: prev date is invalid")
	}

}
