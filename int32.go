package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Int32 SQL type that can retrieve NULL value
type Int32 struct {
	realValue int32
	isValid   bool
}

// NewInt32 creates a new nullable 32-bit integer
func NewInt32(value *int32) Int32 {
	if value == nil {
		return Int32{
			realValue: 0,
			isValid:   false,
		}
	}
	return Int32{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 32-bit integer
func (n Int32) Get() *int32 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 32-bit integer
func (n *Int32) Set(value *int32) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Int32) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed int32
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Int32) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	var i64 int64
	if err := convertAssign(&i64, value); err != nil {
		return err
	}
	n.realValue = int32(i64)

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Int32) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return int64(n.realValue), nil
}

// GormDataType gorm common data type
func (Int32) GormDataType() string {
	return "int32_null"
}

// GormDBDataType gorm db data type
func (Int32) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "INT"
	case "postgres":
		return "integer"
	}
	return ""
}
