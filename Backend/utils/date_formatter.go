package utils

import "time"

type DateFormatter struct {
	Format string
}

func (d *DateFormatter) ParseYearMonthDayOfDateString(date string) (time.Time, error) {
	t, err := time.Parse(d.Format, date)
	if err != nil {
		return time.Time{}, err
	}

	y, m, day := t.Date()
	newDate := time.Date(y, m, day, 0, 0, 0, 0, t.Location())
	return newDate, nil
}
