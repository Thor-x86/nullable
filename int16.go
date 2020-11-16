package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Int16 SQL type that can retrieve NULL value
type Int16 struct {
	realValue int16
	isValid   bool
}

// NewInt16 creates a new nullable 16-bit integer
func NewInt16(value *int16) Int16 {
	if value == nil {
		return Int16{
			realValue: 0,
			isValid:   false,
		}
	}
	return Int16{
		realValue: *value,
		isValid:   true,
	}
}

// Get either nil or 16-bit integer
func (n Int16) Get() *int16 {
	if !n.isValid {
		return nil
	}
	return &n.realValue
}

// Set either nil or 16-bit integer
func (n *Int16) Set(value *int16) {
	n.isValid = (value != nil)
	if n.isValid {
		n.realValue = *value
	} else {
		n.realValue = 0
	}
}

// MarshalJSON converts current value to JSON
func (n Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Int16) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		n.isValid = false
		n.realValue = 0
		return nil
	}

	var parsed int16
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.isValid = true
	n.realValue = parsed
	return nil
}

// Scan implements scanner interface
func (n *Int16) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = 0, false
		return nil
	}

	var i64 int64
	if err := convertAssign(&i64, value); err != nil {
		return err
	}
	n.realValue = int16(i64)

	n.isValid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n Int16) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return int64(n.realValue), nil
}

// GormDataType gorm common data type
func (Int16) GormDataType() string {
	return "int16_null"
}

// GormDBDataType gorm db data type
func (Int16) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "SMALLINT"
	case "postgres":
		return "smallint"
	}
	return ""
}
