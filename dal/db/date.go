package mydb

import "time"

type Date struct {
	Year  int64 `json:year`
	Month int64 `json:month`
	Day   int64 `json:day`
}

func (d *Date) Equals(a *Date) bool {
	return d.Day == a.Day && d.Month == a.Month && d.Year == a.Year
}

func (d *Date) Reset(t time.Time) *Date {
	d.Year = int64(t.Year())
	d.Month = int64(t.Month())
	d.Day = int64(t.Day())
	return d
}
