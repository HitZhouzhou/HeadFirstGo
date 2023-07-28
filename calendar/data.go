package calendar

import "errors"

type Data struct {
	year  int
	month int
	day   int
}

func (d *Data) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year")
	}
	d.year = year
	return nil
}
func (d *Data) SetMonth(month int) error {
	if month > 13 || month < 1 {
		return errors.New("Invalid month")
	}
	d.month = month
	return nil
}
func (d *Data) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.day = day
	return nil
}
func (d *Data) Year() int {
	return d.year
}
func (d *Data) Month() int {
	return d.month
}
func (d *Data) Day() int {
	return d.day
}
