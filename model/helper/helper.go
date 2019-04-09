package helper

import (
	"database/sql"
	"encoding/json"
)

//ToNullString validates a sql.NullString
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

//ToNullInt64 validates a sql.NullInt64
func ToNullInt64(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: true}
}

//ToNullFloat64 validates a sql.NullFloat64
func ToNullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: float64(f), Valid: true}
}

//ToNullFloat64 validates a sql.NullFloat64
func ToNullBool(f bool) sql.NullBool {
	return sql.NullBool{Bool: f, Valid: true}
}

// for JSON to DB & vice-versa
type JsonNullInt64 struct {
	sql.NullInt64
}

// IsEqual will return true if the JsonNullInt64 is valid and equal
// to the given int
func (v JsonNullInt64) IsEqual(i int) bool {
	if v.Valid == true && int(v.Int64) == i {
		return true
	}
	return false
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type JsonNullBool struct {
	sql.NullBool
}

// IsEqual will return true if the JsonNullBool is valid and equal
// to the given boolean
func (v JsonNullBool) IsEqual(b bool) bool {
	if v.Valid == true && v.Bool == b {
		return true
	}
	return false
}

func (v JsonNullBool) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullBool) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Bool = *x
	} else {
		v.Valid = false
	}
	return nil
}

type JsonNullString struct {
	sql.NullString
}

// IsEqual will return true if the JsonNullString is valid and equal
// to the given string
func (v JsonNullString) IsEqual(s string) bool {
	if v.Valid == true && v.String == s {
		return true
	}
	return false
}

func (v JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

type JsonNullFloat64 struct {
	sql.NullFloat64
}

// IsEqual will return true if the JsonNullFloat64 is valid and equal
// to the given float64
func (v JsonNullFloat64) IsEqual(f float64) bool {
	if v.Valid == true && v.Float64 == f {
		return true
	}
	return false
}

func (v JsonNullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullFloat64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *float64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Float64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

//ToJsonNullString validates a JsonNullString
func ToJsonNullString(s string) JsonNullString {
	return JsonNullString{
		sql.NullString{String: s, Valid: true},
	}
}

//ToJsonNullInt64 validates a JsonNullInt64
func ToJsonNullInt64(i int) JsonNullInt64 {
	return JsonNullInt64{
		sql.NullInt64{Int64: int64(i), Valid: true},
	}
}

//ToJsonNullFloat64 validates a JsonNullFloat64
func ToJsonNullFloat64(f float64) JsonNullFloat64 {
	return JsonNullFloat64{
		sql.NullFloat64{Float64: float64(f), Valid: true},
	}
}

//ToJsonNullFloat64 validates a JsonNullBool
func ToJsonNullBool(f bool) JsonNullBool {
	return JsonNullBool{
		sql.NullBool{Bool: f, Valid: true},
	}
}
