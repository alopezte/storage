package storage

import (
	"encoding/json"
	"errors"
	"time"
)

type ShortDate time.Time

func (Date ShortDate) MarshalJSON() ([]byte, error) {
	t := time.Time(Date)
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	return []byte(t.Format(`"02/01/2006"`)), nil
}

func (Date ShortDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, _ := time.Parse("2006-01-02", s)
	Date = ShortDate(t)
	return nil
}
