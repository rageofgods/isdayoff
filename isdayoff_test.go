package isdayoff

import (
	"testing"
	"time"
)

func TestIsLeap(t *testing.T) {
	client := New()
	leap, err := client.IsLeap(2020)
	if err != nil {
		t.Error(err)
	}
	if leap != true {
		t.Errorf("should be true, equal: %v", leap)
	}
}

func TestGetByRange(t *testing.T) {
	client := New()
	countryCode := CountryCodeRussia
	pre := false
	covid := false
	sd := "20210101"
	ed := "20210115"

	day, err := client.GetByRange(ParamsRange{
		StartDate: &sd,
		EndDate:   &ed,
		Params: Params{
			CountryCode: &countryCode,
			Pre:         &pre,
			Covid:       &covid,
		},
	})
	if err != nil {
		t.Error(err)
	}

	const shortForm = "20060102"
	tm, err := time.Parse(shortForm, ed)
	if err != nil {
		t.Errorf("error in date parse: %v", err)
	}

	if len(day) < tm.Day() {
		t.Errorf("should be 15, equal: %v", len(day))
	}

	if day[0] != DayTypeNonWorking {
		t.Errorf("should be 1, equal: %v", day[0])
	}

	if day[14] != DayTypeWorking {
		t.Errorf("should be 0, equal: %v", day[14])
	}
}

func TestGetByYear(t *testing.T) {
	client := New()
	days, err := client.GetBy(Params{Year: 2020})
	if err != nil {
		t.Error(err)
	}
	if len(days) != 366 {
		t.Errorf("should be 366, equal: %v", len(days))
	}
}

func TestGetByDay(t *testing.T) {
	client := New()
	month := time.January
	day := 1
	countryCode := CountryCodeKazakhstan
	pre := false
	covid := false
	days, err := client.GetBy(Params{
		Year:        2020,
		Month:       &month,
		Day:         &day,
		CountryCode: &countryCode,
		Pre:         &pre,
		Covid:       &covid,
	})
	if err != nil {
		t.Error(err)
	}
	if len(days) != 1 {
		t.Errorf("should be 1, equal: %v", len(days))
	}
}

func TestToday(t *testing.T) {
	client := New()
	countryCode := CountryCodeKazakhstan
	pre := false
	covid := false
	day, err := client.Today(Params{
		CountryCode: &countryCode,
		Pre:         &pre,
		Covid:       &covid,
	})
	if err != nil {
		t.Error(err)
	}
	// This is a dynamicly set parameter, so it can vary from day to day.
	if *day != DayTypeNonWorking {
		t.Error("should be non working")
	}
}

func TestTomorrow(t *testing.T) {
	client := New()
	countryCode := CountryCodeKazakhstan
	pre := false
	covid := false
	day, err := client.Tomorrow(Params{
		CountryCode: &countryCode,
		Pre:         &pre,
		Covid:       &covid,
	})
	if err != nil {
		t.Error(err)
	}
	// This is a dynamicly set parameter, so it can vary from day to day.
	if *day != DayTypeWorking {
		t.Errorf("should be %v, instead: %v", DayTypeWorking, *day)
	}
}
