package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Int64 SQL type that can retrieve NULL value
type Int64 struct {
	realValue int64
	isValid   bool
}

// NewInt64 creates a new nullable 64-bit integer
func NewInt64(value *int64) Int64 {
	if value == nil {
		return Int64{
			realValue: 0,
			isValid:   false,
		}
	}
	return Int64{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 64-bit integer
func (n Int64) Get() *int64 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 64-bit integer
func (n *Int64) Set(value *int64) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Int64) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed int64
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Int64) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}
	n.isValid = true
	return convertAssign(&n.realValue, value)
}

// Value implements the driver Valuer interface.
func (n Int64) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return n.realValue, nil
}

// GormDataType gorm common data type
func (Int64) GormDataType() string {
	return "int64_null"
}

// GormDBDataType gorm db data type
func (Int64) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "BIGINT"
	case "postgres":
		return "bigint"
	}
	return ""
}
