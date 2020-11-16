package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Int8 SQL type that can retrieve NULL value
type Int8 struct {
	realValue int8
	isValid   bool
}

// NewInt8 creates a new nullable 8-bit integer
func NewInt8(value *int8) Int8 {
	if value == nil {
		return Int8{
			realValue: 0,
			isValid:   false,
		}
	}
	return Int8{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 8-bit integer
func (n Int8) Get() *int8 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 8-bit integer
func (n *Int8) Set(value *int8) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Int8) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed int8
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Int8) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	var i64 int64
	if err := convertAssign(&i64, value); err != nil {
		return err
	}
	n.realValue = int8(i64)

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Int8) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return int64(n.realValue), nil
}

// GormDataType gorm common data type
func (Int8) GormDataType() string {
	return "int8_null"
}

// GormDBDataType gorm db data type
func (Int8) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "TINYINT"
	case "postgres":
		return "smallint"
	}
	return ""
}
