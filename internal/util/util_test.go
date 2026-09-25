package util

import (
	"testing"
	"time"
)

func TestNearestSeason(t *testing.T) {
	cases := []struct {
		date   string
		year   int
		season int
	}{
		{"2026-01-01", 2026, 1},
		{"2026-02-28", 2026, 1},
		{"2026-03-01", 2026, 4},
		{"2026-04-15", 2026, 4},
		{"2026-05-31", 2026, 4},
		{"2026-06-01", 2026, 7},
		{"2026-07-01", 2026, 7},
		{"2026-08-01", 2026, 7},
		{"2026-09-25", 2026, 10},
		{"2026-10-01", 2026, 10},
		{"2026-11-30", 2026, 10},
		{"2026-12-01", 2027, 1},
		{"2026-12-31", 2027, 1},
	}

	for _, c := range cases {
		d, err := time.Parse("2006-01-02", c.date)
		if err != nil {
			t.Fatalf("parse %s: %v", c.date, err)
		}
		y, s := NearestSeason(d)
		if y != c.year || s != c.season {
			t.Errorf("NearestSeason(%s) = %04d/%02d, want %04d/%02d", c.date, y, s, c.year, c.season)
		}
	}
}
