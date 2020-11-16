package nullable

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Time SQL type that can retrieve NULL value
type Time struct {
	realValue time.Time
	isValid   bool
}

// NewTime creates a new nullable 64-bit integer
func NewTime(value *time.Time) Time {
	if value == nil {
		return Time{
			realValue: time.Time{},
			isValid:   false,
		}
	}
	return Time{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 64-bit integer
func (n Time) Get() *time.Time {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 64-bit integer
func (n *Time) Set(value *time.Time) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = time.Time{}
	}
}

// MarshalJSON converts current value to JSON
func (n Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Time) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = time.Time{}
		return nil
	}

	var parsed time.Time
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Time) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = time.Time{}, false
		return nil
	}

	var utcTime time.Time
	if err := convertAssign(&utcTime, value); err != nil {
		return err
	}
	n.realValue = utcTime.Local()

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Time) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return n.realValue.UTC(), nil
}

// GormDataType gorm common data type
func (Time) GormDataType() string {
	return "timestamp_null"
}

// GormDBDataType gorm db data type
func (Time) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "DATETIME"
	case "mysql":
		return "TIMESTAMP NULL DEFAULT NULL"
	case "postgres":
		return "timestamp"
	}
	return ""
}
