package timepack

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

var DefaultLayout = "2006-01-02 15:04:05"

type PackNull struct {
	Layout string
	time.Time
}

func NewNull(t time.Time) PackNull {
	return PackNull{
		Layout: DefaultLayout,
		Time:   t,
	}
}

func (p *PackNull) Scan(value interface{}) error {
	if value == nil {
		p.Time = time.Time{}
		return nil
	}
	p.Time = value.(time.Time)
	return nil
}

func (p PackNull) Value() (driver.Value, error) {
	if p.IsZero() {
		return nil, nil
	}
	return p.Time, nil
}

func (p *PackNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		p.Time = time.Time{}
		return nil
	}
	if p.Layout == "" {
		p.Layout = DefaultLayout
	}
	newt, err := time.Parse(p.Layout, strings.Replace(string(data), `"`, "", -1))
	if err != nil {
		p.Time = time.Time{}
		return err
	}
	p.Time = newt
	return nil
}

func (p *PackNull) UnmarshalText(text []byte) error {
	return p.UnmarshalJSON(text)
}

func (p PackNull) MarshalJSON() ([]byte, error) {
	if p.IsZero() {
		return []byte("null"), nil
	}
	if p.Layout == "" {
		p.Layout = DefaultLayout
	}
	str := p.Time.Format(p.Layout)
	return json.Marshal(str)
}

func (p PackNull) MarshalText() ([]byte, error) {
	if p.IsZero() {
		return []byte{}, nil
	}
	if p.Layout == "" {
		p.Layout = DefaultLayout
	}
	return []byte(p.Time.Format(p.Layout)), nil
}

type PackZero struct {
	Layout string
	time.Time
}

func NewZero(t time.Time) PackZero {
	return PackZero{
		Layout: DefaultLayout,
		Time:   t,
	}
}

func (p *PackZero) Scan(value interface{}) error {
	if value == nil {
		p.Time = time.Time{}
		return nil
	}
	p.Time = value.(time.Time)
	return nil
}

func (p PackZero) Value() (driver.Value, error) {
	if p.IsZero() {
		return time.Time{}, nil
	}
	return p.Time, nil
}

func (p *PackZero) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		p.Time = time.Time{}
		return nil
	}
	if p.Layout == "" {
		p.Layout = DefaultLayout
	}
	newt, err := time.Parse(p.Layout, strings.Replace(string(data), `"`, "", -1))
	if err != nil {
		p.Time = time.Time{}
		return err
	}
	p.Time = newt
	return nil
}

func (p *PackZero) UnmarshalText(text []byte) error {
	return p.UnmarshalJSON(text)
}

func (p PackZero) MarshalJSON() ([]byte, error) {
	if p.IsZero() {
		return []byte(""), nil
	}
	if p.Layout == "" {
		p.Layout = DefaultLayout
	}
	str := p.Time.Format(p.Layout)
	return json.Marshal(str)
}

func (p PackZero) MarshalText() ([]byte, error) {
	if p.IsZero() {
		return []byte{}, nil
	}
	if p.Layout == "" {
		p.Layout = DefaultLayout
	}
	return []byte(p.Time.Format(p.Layout)), nil
}
