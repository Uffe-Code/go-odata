package date

import (
	"encoding/json"
	"time"
)

// Date represents a date without a time or timezone
type Date struct {
	time time.Time
}

const (
	RFC3339 = "2006-01-02"
)

// Today gets current day
func Today() Date {
	return Date{time.Now()}
}

func New(year int, month time.Month, day int) Date {
	return Date{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func Parse(layout string, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	return Date{t}, err
}

func (date Date) Time() time.Time {
	t, _ := time.Parse(RFC3339, date.String())
	return t
}

func (date Date) String() string {
	return date.time.Format(RFC3339)
}

func (date Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(date.String())
}

func (date *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	*date, err = Parse(`"`+RFC3339+`"`, string(data))
	return err
}
