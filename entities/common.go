package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	var s *time.Time
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Time = *s
	} else {
		v.Valid = false
	}
	return nil
}

type NullString struct {
	sql.NullString
}

func (nt *NullString) Scan(value interface{}) error {
	switch value.(type) {
	case []byte:
		nt.String, nt.Valid = string(value.([]byte)), true
	case string:
		nt.String, nt.Valid = value.(string)
	}
	return nil
}

func (nt NullString) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.String, nil
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.String = *s
	} else {
		v.Valid = false
	}
	return nil
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nt *NullFloat64) Scan(value interface{}) error {
	nt.Float64, nt.Valid = value.(float64)
	return nil
}

func (nt NullFloat64) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Float64, nil
}

func (v NullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	}
	return json.Marshal(nil)

}

func (v *NullFloat64) UnmarshalJSON(data []byte) error {
	var s *float64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Float64 = *s
	} else {
		v.Valid = false
	}
	return nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (nt *NullInt64) Scan(value interface{}) error {
	nt.Int64, nt.Valid = value.(int64)
	return nil
}

func (nt NullInt64) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Int64, nil
}

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)

}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	var s *int64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Int64 = *s
	} else {
		v.Valid = false
	}
	return nil
}
