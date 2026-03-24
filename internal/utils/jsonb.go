package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONB represents a JSONB column in PostgreSQL
type JSONB map[string]interface{}

// Value implements the driver.Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan implements the sql.Scanner interface
func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

// GetString retrieves a string value from JSONB
func (j JSONB) GetString(key string) (string, bool) {
	if v, ok := j[key]; ok {
		if s, ok := v.(string); ok {
			return s, true
		}
	}
	return "", false
}

// GetInt retrieves an integer value from JSONB
func (j JSONB) GetInt(key string) (int, bool) {
	if v, ok := j[key]; ok {
		switch val := v.(type) {
		case int:
			return val, true
		case int64:
			return int(val), true
		case float64:
			return int(val), true
		}
	}
	return 0, false
}

// GetFloat retrieves a float64 value from JSONB
func (j JSONB) GetFloat(key string) (float64, bool) {
	if v, ok := j[key]; ok {
		if f, ok := v.(float64); ok {
			return f, true
		}
	}
	return 0, false
}

// GetBool retrieves a boolean value from JSONB
func (j JSONB) GetBool(key string) (bool, bool) {
	if v, ok := j[key]; ok {
		if b, ok := v.(bool); ok {
			return b, true
		}
	}
	return false, false
}

// GetJSONB retrieves a nested JSONB value from JSONB
func (j JSONB) GetJSONB(key string) (JSONB, bool) {
	if v, ok := j[key]; ok {
		if m, ok := v.(map[string]interface{}); ok {
			return JSONB(m), true
		}
	}
	return nil, false
}

// Set sets a value in JSONB
func (j *JSONB) Set(key string, value interface{}) {
	if *j == nil {
		*j = make(JSONB)
	}
	(*j)[key] = value
}

// IsEmpty checks if JSONB is empty
func (j JSONB) IsEmpty() bool {
	return len(j) == 0
}

// ToJSON converts JSONB to JSON string
func (j JSONB) ToJSON() (string, error) {
	bytes, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
