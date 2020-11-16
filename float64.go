package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Float64 SQL type that can retrieve NULL value
type Float64 struct {
	realValue float64
	isValid   bool
}

// NewFloat64 creates a new nullable double precision float
func NewFloat64(value *float64) Float64 {
	if value == nil {
		return Float64{
			realValue: 0,
			isValid:   false,
		}
	}
	return Float64{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or double precision float
func (n Float64) Get() *float64 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or double precision float
func (n *Float64) Set(value *float64) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Float64) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed float64
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Float64) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}
	n.isValid = true
	return convertAssign(&n.realValue, value)
}

// Value implements the driver Valuer interface.
func (n Float64) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return n.realValue, nil
}

// GormDataType gorm common data type
func (Float64) GormDataType() string {
	return "float64_null"
}

// GormDBDataType gorm db data type
func (Float64) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "DOUBLE"
	case "postgres":
		return "double precision"
	}
	return ""
}
