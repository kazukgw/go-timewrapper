package timepack

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

var Format = "2006-01-02 15:04:05"

type Pack struct {
	time.Time
}

func (p *Pack) Scan(value interface{}) error {
	if value == nil {
		p.Time = time.Time{}
		return nil
	}
	p.Time = value.(time.Time)
	return nil
}

func (p Pack) Value() (driver.Value, error) {
	if p.IsZero() {
		return nil, nil
	}
	return p.Time, nil
}

func (p *Pack) UnmarshalJSON(data []byte) error {
	newt, err := time.Parse(Format, strings.Replace(string(data), `"`, "", -1))
	if err != nil {
		p.Time = time.Time{}
		return err
	}
	p.Time = newt
	return nil
}

func (p *Pack) UnmarshalText(text []byte) error {
	return p.UnmarshalJSON(text)
}

func (p Pack) MarshalJSON() ([]byte, error) {
	if p.IsZero() {
		return []byte("null"), nil
	}
	str := p.Time.Format(Format)
	return json.Marshal(str)
}

func (p Pack) MarshalText() ([]byte, error) {
	if p.IsZero() {
		return []byte{}, nil
	}
	return []byte(p.Time.Format(Format)), nil
}
